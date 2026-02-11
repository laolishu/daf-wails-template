import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      wailsjs: path.resolve(__dirname, 'wailsjs'),
    },
    // 确保 React 只有一个实例，避免 useId 等 hooks 问题
    dedupe: ['react', 'react-dom'],
  },
  optimizeDeps: {
    include: ['react', 'react-dom', '@mantine/core', '@mantine/hooks'],
  },
  server: {
    host: 'localhost',
    port: 5173,
    strictPort: true,
  },
})
