import { Button, Stack, Text, Title } from "@mantine/core";

type UpgradeMessages = {
  upgradeTitle: string;
  upgradeDesc: string;
  upgradeCta: string;
};

type UpgradePageProps = {
  messages: UpgradeMessages;
};

export default function UpgradePage({ messages }: UpgradePageProps) {
  return (
    <Stack gap="sm" w="100%">
      <Title order={2} c="#1F2937">{messages.upgradeTitle}</Title>
      <Text c="#6B7280">{messages.upgradeDesc}</Text>
      <Button
        size="md"
        radius={22}
        styles={{
          root: {
            backgroundColor: "#3B82F6",
            color: "#FFFFFF",
            borderRadius: 22,
            "&:hover": {
              backgroundColor: "#2563EB",
            },
            "&:active": {
              backgroundColor: "#1D4ED8",
            },
          },
        }}
      >
        {messages.upgradeCta}
      </Button>
    </Stack>
  );
}
