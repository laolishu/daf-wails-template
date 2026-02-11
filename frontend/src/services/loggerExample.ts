/*
 * @Descripttion: 
 * @version: 
 * @Author: lfzxs@qq.com
 * @Date: 2026-02-10 23:17:30
 * @LastEditors: lfzxs@qq.com
 * @LastEditTime: 2026-02-10 23:29:45
 */
import { WriteTestLog } from "wailsjs/go/backend/App";
import type { LogWriteResult } from "../types/logWriteResult";
import { LogInfo } from "wailsjs/runtime/runtime";

export async function writeTestLog(): Promise<LogWriteResult> {
  LogInfo("Calling WriteTestLog from frontend...");
  return WriteTestLog();
}
