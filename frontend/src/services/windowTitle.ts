import { GetWindowTitle } from "wailsjs/go/backend/App";

export async function fetchWindowTitle(): Promise<string> {
  return GetWindowTitle();
}
