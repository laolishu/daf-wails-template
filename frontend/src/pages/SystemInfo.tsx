import { Alert, Group, Stack, Text, Title } from "@mantine/core";
import { useEffect, useState } from "react";
import { fetchSystemInfo } from "../services/systemInfo";
import type { SystemInfo } from "../types/systemInfo";

export default function SystemInfoPage() {
  const [info, setInfo] = useState<SystemInfo | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchSystemInfo()
      .then((data) => setInfo(data))
      .catch((err: Error) => setError(err.message));
  }, []);

  return (
    <Stack gap="md" w="100%">
      <Title order={2} c="#1F2937">About / System Info</Title>
      <Text c="#6B7280">
        Read-only system details for support and diagnostics.
      </Text>
      {error && (
        <Alert color="red" title="System Info Error">
          {error}
        </Alert>
      )}
      {info && (
        <Stack gap="xs">
          <Group justify="space-between">
            <Text c="#6B7280">Version</Text>
            <Text c="#4B5563">{info.version}</Text>
          </Group>
          <Group justify="space-between">
            <Text c="#6B7280">Build Time</Text>
            <Text c="#4B5563">{info.buildTime}</Text>
          </Group>
          <Group justify="space-between">
            <Text c="#6B7280">Environment</Text>
            <Text c="#4B5563">{info.environment}</Text>
          </Group>
        </Stack>
      )}
    </Stack>
  );
}
