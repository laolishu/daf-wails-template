import {
  Box,
  Button,
  Stack,
  Text,
  Title,
  ThemeIcon,
  useMantineTheme,
} from "@mantine/core";
import { IconHeartbeat } from "@tabler/icons-react";

type HealthMessages = {
  healthTitle: string;
  healthDesc: string;
  healthCta: string;
};

type HealthCheckProps = {
  messages: HealthMessages;
};

export default function HealthCheck({ messages }: HealthCheckProps) {
  const theme = useMantineTheme();

  return (
    <Stack gap="md" w="100%" align="center" ta="center" py="xl">
      <Box
        w={220}
        h={220}
        style={{
          borderRadius: theme.radius.xl,
          backgroundColor: "#EEF2F7",
          border: "1px solid #E2E8F0",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
        }}
      >
        <ThemeIcon
          size={72}
          radius="xl"
          styles={{
            root: {
              backgroundColor: "#E8EEF8",
              color: "#3B82F6",
            },
          }}
        >
          <IconHeartbeat size={36} />
        </ThemeIcon>
      </Box>
      <Title order={2} c="#1F2937">
        {messages.healthTitle}
      </Title>
      <Text c="#6B7280">{messages.healthDesc}</Text>
      <Button
        size="md"
        radius={22}
        px={28}
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
        {messages.healthCta}
      </Button>
    </Stack>
  );
}
