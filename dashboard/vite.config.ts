import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import vercel from 'vite-plugin-vercel';

export default defineConfig({
  plugins: [react(), vercel()],
  optimizeDeps: {
    exclude: ['lucide-react']
  },
  server: {
    proxy: {
      '/events': {
        target: 'http://localhost:3322',
        changeOrigin: true
      }
    }
  }
});