import { CheckForUpdate, DownloadUpdate } from "wailsjs/go/backend/App";
import type { UpdateCheckResponse } from "../types/updateCheckResponse";
import type { UpdateInfo } from "../types/updateInfo";
import type { UpdateInstallResult } from "../types/updateInstallResult";

export async function checkForUpdate(channel: string): Promise<UpdateCheckResponse> {
  return CheckForUpdate(channel);
}

export async function downloadUpdate(info: UpdateInfo): Promise<UpdateInstallResult> {
  const result = await DownloadUpdate(info);
  return { path: result.Path };
}
