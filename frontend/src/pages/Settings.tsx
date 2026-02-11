import { Alert, Divider, Group, Select, Stack, Text, Title } from "@mantine/core";
import { useEffect, useState } from "react";
import { fetchSystemInfo } from "../services/systemInfo";
import type { SystemInfo } from "../types/systemInfo";
import type { Language } from "../i18n";

type SettingsMessages = {
  settingsTitle: string;
  settingsDesc: string;
  settingsGeneralSection: string;
  settingsSystemSection: string;
  settingsLanguageLabel: string;
  settingsLanguageZh: string;
  settingsLanguageEn: string;
  settingsSystemInfoTitle: string;
  settingsSystemInfoDesc: string;
  settingsSystemInfoErrorTitle: string;
  systemVersionLabel: string;
  systemBuildTimeLabel: string;
  systemEnvironmentLabel: string;
};

type SettingsPageProps = {
  language: Language;
  messages: SettingsMessages;
  onLanguageChange: (value: Language) => void;
};

export default function SettingsPage({
  language,
  messages,
  onLanguageChange,
}: SettingsPageProps) {
  const [info, setInfo] = useState<SystemInfo | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchSystemInfo()
      .then((data) => setInfo(data))
      .catch((err: Error) => setError(err.message));
  }, []);

  return (
    <Stack gap="md" w="100%">
      <Stack gap="sm">
        <Title order={2} c="#1F2937">
          {messages.settingsTitle}
        </Title>
        <Text c="#6B7280">{messages.settingsDesc}</Text>
      </Stack>

      <Stack gap="xs">
        <Title order={3} c="#1F2937">
          {messages.settingsGeneralSection}
        </Title>
        <Group justify="space-between" align="center" wrap="nowrap">
          <Text c="#4B5563" fw={500}>
            {messages.settingsLanguageLabel}
          </Text>
          <Select
            value={language}
            data={[
              { value: "zh-CN", label: messages.settingsLanguageZh },
              { value: "en-US", label: messages.settingsLanguageEn },
            ]}
            onChange={(value) => {
              if (value === "zh-CN" || value === "en-US") {
                onLanguageChange(value);
              }
            }}
          />
        </Group>
      </Stack>

      <Divider />

      <Stack gap="sm">
        <Title order={3} c="#1F2937">
          {messages.settingsSystemSection}
        </Title>
        <Text c="#6B7280">{messages.settingsSystemInfoDesc}</Text>
        {error && (
          <Alert color="red" title={messages.settingsSystemInfoErrorTitle}>
            {error}
          </Alert>
        )}
        {info && (
          <Stack gap="xs">
            <Group justify="space-between">
              <Text c="#6B7280">{messages.systemVersionLabel}</Text>
              <Text c="#4B5563">{info.version}</Text>
            </Group>
            <Group justify="space-between">
              <Text c="#6B7280">{messages.systemBuildTimeLabel}</Text>
              <Text c="#4B5563">{info.buildTime}</Text>
            </Group>
            <Group justify="space-between">
              <Text c="#6B7280">{messages.systemEnvironmentLabel}</Text>
              <Text c="#4B5563">{info.environment}</Text>
            </Group>
          </Stack>
        )}
      </Stack>
    </Stack>
  );
}
