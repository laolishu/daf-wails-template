import { Alert, Badge, Button, Group, Modal, Stack, Text, Title } from "@mantine/core";
import { useState } from "react";
import { checkForUpdate, downloadUpdate } from "../services/updater";
import type { UpdateCheckResponse } from "../types/updateCheckResponse";

type UpgradeMessages = {
  upgradeTitle: string;
  upgradeDesc: string;
  upgradeCta: string;
  upgradeModalTitle: string;
  upgradeChecking: string;
  upgradeCheckError: string;
  upgradeCurrentLabel: string;
  upgradeLatestLabel: string;
  upgradeForceLabel: string;
  upgradeForceYes: string;
  upgradeForceNo: string;
  upgradeReleaseNotesLabel: string;
  upgradeDownloadUrlLabel: string;
  upgradeDownloadButton: string;
  upgradeCancelButton: string;
  upgradeDownloadResultTitle: string;
  upgradeDownloadSuccess: string;
  upgradeDownloadError: string;
};

type UpgradePageProps = {
  messages: UpgradeMessages;
};

export default function UpgradePage({ messages }: UpgradePageProps) {
  const [checking, setChecking] = useState(false);
  const [downloadLoading, setDownloadLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [checkResult, setCheckResult] = useState<UpdateCheckResponse | null>(null);
  const [modalOpen, setModalOpen] = useState(false);
  const [downloadPath, setDownloadPath] = useState<string | null>(null);

  const handleCheckUpdate = async () => {
    setChecking(true);
    setError(null);
    setDownloadPath(null);
    try {
      const result = await checkForUpdate("stable");
      setCheckResult(result);
      setModalOpen(true);
    } catch (err) {
      const message = err instanceof Error ? err.message : messages.upgradeCheckError;
      setError(message);
    } finally {
      setChecking(false);
    }
  };

  const handleDownload = async () => {
    if (!checkResult) {
      return;
    }
    setDownloadLoading(true);
    setError(null);
    try {
      const result = await downloadUpdate(checkResult.info);
      setDownloadPath(result.path);
    } catch (err) {
      const message = err instanceof Error ? err.message : messages.upgradeDownloadError;
      setError(message);
    } finally {
      setDownloadLoading(false);
    }
  };

  return (
    <Stack gap="sm" w="100%">
      <Title order={2} c="#1F2937">{messages.upgradeTitle}</Title>
      <Text c="#6B7280">{messages.upgradeDesc}</Text>
      {error && (
        <Alert color="red" title={messages.upgradeCheckError}>
          {error}
        </Alert>
      )}
      <Button
        size="md"
        radius={22}
        loading={checking}
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
        onClick={handleCheckUpdate}
      >
        {messages.upgradeCta}
      </Button>

      <Modal
        opened={modalOpen}
        onClose={() => setModalOpen(false)}
        title={messages.upgradeModalTitle}
        centered
      >
        {checking && <Text c="#6B7280">{messages.upgradeChecking}</Text>}
        {checkResult && (
          <Stack gap="xs">
            <Group justify="space-between">
              <Text c="#6B7280">{messages.upgradeCurrentLabel}</Text>
              <Text c="#111827">{checkResult.currentVersion || "-"}</Text>
            </Group>
            <Group justify="space-between">
              <Text c="#6B7280">{messages.upgradeLatestLabel}</Text>
              <Text c="#111827">{checkResult.info.latestVersion || "-"}</Text>
            </Group>
            <Group justify="space-between">
              <Text c="#6B7280">{messages.upgradeForceLabel}</Text>
              <Badge color={checkResult.info.force ? "red" : "gray"} variant="light">
                {checkResult.info.force
                  ? messages.upgradeForceYes
                  : messages.upgradeForceNo}
              </Badge>
            </Group>
            {checkResult.info.releaseNotes && (
              <Stack gap={4}>
                <Text c="#6B7280">{messages.upgradeReleaseNotesLabel}</Text>
                <Text c="#111827" style={{ whiteSpace: "pre-wrap" }}>
                  {checkResult.info.releaseNotes}
                </Text>
              </Stack>
            )}
            {checkResult.info.downloadUrl && (
              <Stack gap={4}>
                <Text c="#6B7280">{messages.upgradeDownloadUrlLabel}</Text>
                <Text c="#111827" style={{ wordBreak: "break-all" }}>
                  {checkResult.info.downloadUrl}
                </Text>
              </Stack>
            )}
            {downloadPath && (
              <Alert color="green" title={messages.upgradeDownloadResultTitle}>
                {messages.upgradeDownloadSuccess} {downloadPath}
              </Alert>
            )}
            <Group justify="flex-end" mt="sm">
              <Button variant="default" onClick={() => setModalOpen(false)}>
                {messages.upgradeCancelButton}
              </Button>
              <Button loading={downloadLoading} onClick={handleDownload}>
                {messages.upgradeDownloadButton}
              </Button>
            </Group>
          </Stack>
        )}
      </Modal>
    </Stack>
  );
}
