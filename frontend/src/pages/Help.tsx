import { Stack, Text, Title } from "@mantine/core";

type HelpMessages = {
  helpTitle: string;
  helpDesc: string;
};

type HelpPageProps = {
  messages: HelpMessages;
};

export default function HelpPage({ messages }: HelpPageProps) {
  return (
    <Stack gap="sm" w="100%">
      <Title order={2} c="#1F2937">{messages.helpTitle}</Title>
      <Text c="#6B7280">{messages.helpDesc}</Text>
    </Stack>
  );
}
