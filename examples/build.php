<?php
/**
 * Compile a sketch through the Mujard Developer API: submit, poll, download.
 *
 *   MUJARD_KEY=mjd_test_… php examples/build.php
 *
 * Optional:
 *   MUJARD_BASE=http://localhost/mujard-api/api   (default: the hosted API)
 *   MUJARD_BOARD=uno                               (any id from GET /v1/boards, or an FQBN)
 *
 * Keep the key on your server only - never inside an app you ship.
 */

$BASE = rtrim(getenv('MUJARD_BASE') ?: 'https://api.mujard.example', '/');
$KEY  = getenv('MUJARD_KEY') ?: exit("Set MUJARD_KEY=mujard_… first.\n");
$BOARD = getenv('MUJARD_BOARD') ?: 'uno';

/** One JSON call. Returns [http_status, raw_body]. */
function api(string $method, string $path, ?array $body, string $key, array $extraHeaders = []): array
{
    $ch = curl_init($GLOBALS['BASE'] . $path);
    $h = array_merge(['Authorization: Bearer ' . $key, 'Accept: application/json'], $extraHeaders);
    if ($body !== null) {
        $h[] = 'Content-Type: application/json';
        curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($body));
    }
    curl_setopt_array($ch, [CURLOPT_CUSTOMREQUEST => $method, CURLOPT_RETURNTRANSFER => true, CURLOPT_HTTPHEADER => $h, CURLOPT_TIMEOUT => 30]);
    $raw = curl_exec($ch);
    $code = curl_getinfo($ch, CURLINFO_RESPONSE_CODE);
    curl_close($ch);
    if ($raw === false) {
        exit("Could not reach the API. Check MUJARD_BASE.\n");
    }
    return [$code, $raw];
}

/** Turn an error response into a readable message (switch on ->code, never the text). */
function apiFail(int $code, string $raw): string
{
    $e = json_decode($raw, true)['error'] ?? null;
    $where = is_array($e) && isset($e['request_id']) ? " (request {$e['request_id']})" : '';
    return is_array($e) ? "{$e['code']}: {$e['message']}$where" : "HTTP $code$where";
}

// 1. Submit. Idempotency-Key is required: reuse it on a retry and the same build comes back
//    instead of a duplicate charge (the response then carries Idempotent-Replayed: true).
$files = [
    'sketch.ino' => "void setup() { pinMode(LED_BUILTIN, OUTPUT); }\n"
        . "void loop() { digitalWrite(LED_BUILTIN, !digitalRead(LED_BUILTIN)); delay(500); }\n",
];
$idem = 'example-' . bin2hex(random_bytes(6));

[$code, $raw] = api('POST', '/v1/builds', ['board' => $BOARD, 'files' => $files], $KEY, ["Idempotency-Key: $idem"]);

if ($code === 401) {
    exit('The key was refused - ' . apiFail($code, $raw) . "\n");
}
if ($code === 422) {
    exit('The request was rejected - ' . apiFail($code, $raw) . "\n");
}
if ($code === 429) {
    exit("Rate limited - wait a moment and retry with the SAME idempotency key ($idem).\n");
}
if (!in_array($code, [200, 202], true)) {
    exit('Unexpected answer - ' . apiFail($code, $raw) . "\n");
}

$job = json_decode($raw, true);
echo "build {$job['id']}: {$job['state']} ({$job['mode']} mode, {$job['cost']} coins)\n";

// 2. A test key answers instantly; a live key queues for the real toolchain. Poll until settled.
$waited = 0;
while (in_array($job['state'], ['queued', 'running'], true)) {
    if (++$waited > 100) {
        exit("Still building after 5 minutes - check GET /v1/builds/{$job['id']} later.\n");
    }
    sleep(3);
    [$code, $raw] = api('GET', '/v1/builds/' . $job['id'], null, $KEY);
    $job = json_decode($raw, true) ?? $job;
    echo "… {$job['state']}\n";
}

// 3. A successful build carries a firmware artifact - raw bytes, not JSON.
if ($job['state'] === 'done' && ($job['result']['success'] ?? false)) {
    [$code, $bin] = api('GET', '/v1/builds/' . $job['id'] . '/artifact', null, $KEY);
    if ($code !== 200) {
        exit('The build finished but the artifact did not download - ' . apiFail($code, $bin) . "\n");
    }
    $format = $job['result']['format'] ?? 'hex';
    $out = getcwd() . '/firmware.' . $format;
    file_put_contents($out, $bin);
    $s = $job['result']['sizes'] ?? null;
    $sizes = $s ? " ({$s['program_bytes']} bytes program, {$s['ram_bytes']} bytes RAM)" : '';
    echo "saved $out" . $sizes . "\n";
    if (($job['result']['test'] ?? false) === true) {
        echo "note: this was a TEST-mode build - the file is the sample placeholder. Use a live key for real firmware.\n";
    }
} else {
    echo "build {$job['state']}\n";
    foreach ($job['result']['errors'] ?? [] as $e) {
        echo "  $e\n";
    }
}
