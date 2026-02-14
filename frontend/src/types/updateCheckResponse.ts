import type { UpdateInfo } from "./updateInfo";

export type UpdateCheckResponse = {
  info: UpdateInfo;
  currentVersion: string;
};
