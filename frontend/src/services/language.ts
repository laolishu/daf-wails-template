import { GetLanguage, SetLanguage } from "wailsjs/go/backend/App";

export async function fetchLanguage(): Promise<string> {
  return GetLanguage();
}

export async function updateLanguage(language: "zh-CN" | "en-US"): Promise<void> {
  return SetLanguage(language);
}
