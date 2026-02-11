import {
  Badge,
  Button,
  Card,
  Group,
  Stack,
  Text,
  Title,
} from "@mantine/core";
import { useState } from "react";
import { fetchSystemInfo } from "../services/systemInfo";
import { fetchConfigSummary } from "../services/configSummary";
import { writeTestLog } from "../services/loggerExample";
import type { SystemInfo } from "../types/systemInfo";
import type { ConfigSummary } from "../types/configSummary";
import type { LogWriteResult } from "../types/logWriteResult";

type ExamplesMessages = {
  examplesTitle: string;
  examplesDesc: string;
  loggerExampleTitle: string;
  loggerExampleDesc: string;
  loggerExampleButton: string;
  loggerExampleOk: string;
  loggerExampleError: string;
  loggerExampleLogDir: string;
  systemInfoExampleTitle: string;
  systemInfoExampleDesc: string;
  systemInfoExampleButton: string;
  configExampleTitle: string;
  configExampleDesc: string;
  configExampleButton: string;
  systemVersionLabel: string;
  systemBuildTimeLabel: string;
  systemEnvironmentLabel: string;
  configLanguageLabel: string;
  configLogLevelLabel: string;
  configLogDirLabel: string;
};

type ExamplesPageProps = {
  messages: ExamplesMessages;
};

export default function ExamplesPage({ messages }: ExamplesPageProps) {
  const [logResult, setLogResult] = useState<LogWriteResult | null>(null);
  const [logLoading, setLogLoading] = useState(false);
  const [logLocked, setLogLocked] = useState(false);
  const [systemInfo, setSystemInfo] = useState<SystemInfo | null>(null);
  const [systemLoading, setSystemLoading] = useState(false);
  const [configSummary, setConfigSummary] = useState<ConfigSummary | null>(null);
  const [configLoading, setConfigLoading] = useState(false);

  const handleWriteLog = async () => {
    if (logLocked) {
      return;
    }
    setLogLoading(true);
    try {
      const result = await writeTestLog();
      setLogResult(result);
      if (result.ok) {
        setLogLocked(true);
      }
    } catch (err) {
      setLogResult({ ok: false, error: (err as Error).message });
    } finally {
      setLogLoading(false);
    }
  };

  const handleLoadSystemInfo = async () => {
    setSystemLoading(true);
    try {
      const result = await fetchSystemInfo();
      setSystemInfo(result);
    } finally {
      setSystemLoading(false);
    }
  };

  const handleLoadConfig = async () => {
    setConfigLoading(true);
    try {
      const result = await fetchConfigSummary();
      setConfigSummary(result);
    } finally {
      setConfigLoading(false);
    }
  };

  return (
    <Stack gap="md" w="100%">
      <Stack gap="xs">
        <Title order={2} c="#1F2937">
          {messages.examplesTitle}
        </Title>
        <Text c="#6B7280">{messages.examplesDesc}</Text>
      </Stack>

      <Card withBorder radius="md" padding="md">
        <Stack gap="sm">
          <Group justify="space-between">
            <Title order={4}>{messages.loggerExampleTitle}</Title>
            {logResult && (
              <Badge color={logResult.ok ? "green" : "red"} variant="light">
                {logResult.ok ? "OK" : "Error"}
              </Badge>
            )}
          </Group>
          <Text c="#6B7280">{messages.loggerExampleDesc}</Text>
          <Group justify="space-between">
            <Button
              size="sm"
              onClick={handleWriteLog}
              loading={logLoading}
              disabled={logLocked}
            >
              {messages.loggerExampleButton}
            </Button>
            {logResult && logResult.ok && logResult.logDir && (
              <Text c="#4B5563" size="sm">
                {messages.loggerExampleLogDir} {logResult.logDir}
              </Text>
            )}
          </Group>
          {logResult && !logResult.ok && (
            <Text c="red" size="sm">
              {messages.loggerExampleError}: {logResult.error}
            </Text>
          )}
          {logResult && logResult.ok && (
            <Text c="#4B5563" size="sm">
              {messages.loggerExampleOk}
            </Text>
          )}
        </Stack>
      </Card>

      <Card withBorder radius="md" padding="md">
        <Stack gap="sm">
          <Group justify="space-between">
            <Title order={4}>{messages.systemInfoExampleTitle}</Title>
            {systemInfo && (
              <Badge color="blue" variant="light">
                OK
              </Badge>
            )}
          </Group>
          <Text c="#6B7280">{messages.systemInfoExampleDesc}</Text>
          <Button size="sm" onClick={handleLoadSystemInfo} loading={systemLoading}>
            {messages.systemInfoExampleButton}
          </Button>
          {systemInfo && (
            <Stack gap={6}>
              <Group justify="space-between">
                <Text c="#6B7280">{messages.systemVersionLabel}</Text>
                <Text c="#4B5563">{systemInfo.version}</Text>
              </Group>
              <Group justify="space-between">
                <Text c="#6B7280">{messages.systemBuildTimeLabel}</Text>
                <Text c="#4B5563">{systemInfo.buildTime}</Text>
              </Group>
              <Group justify="space-between">
                <Text c="#6B7280">{messages.systemEnvironmentLabel}</Text>
                <Text c="#4B5563">{systemInfo.environment}</Text>
              </Group>
            </Stack>
          )}
        </Stack>
      </Card>

      <Card withBorder radius="md" padding="md">
        <Stack gap="sm">
          <Group justify="space-between">
            <Title order={4}>{messages.configExampleTitle}</Title>
            {configSummary && (
              <Badge color="blue" variant="light">
                OK
              </Badge>
            )}
          </Group>
          <Text c="#6B7280">{messages.configExampleDesc}</Text>
          <Button size="sm" onClick={handleLoadConfig} loading={configLoading}>
            {messages.configExampleButton}
          </Button>
          {configSummary && (
            <Stack gap={6}>
              <Group justify="space-between">
                <Text c="#6B7280">{messages.configLanguageLabel}</Text>
                <Text c="#4B5563">{configSummary.language}</Text>
              </Group>
              <Group justify="space-between">
                <Text c="#6B7280">{messages.configLogLevelLabel}</Text>
                <Text c="#4B5563">{configSummary.logLevel}</Text>
              </Group>
              <Group justify="space-between">
                <Text c="#6B7280">{messages.configLogDirLabel}</Text>
                <Text c="#4B5563">{configSummary.logDir}</Text>
              </Group>
            </Stack>
          )}
        </Stack>
      </Card>
    </Stack>
  );
}
