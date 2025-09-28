import { defineConfig } from 'vite';
import tailwindcss from '@tailwindcss/vite'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        tailwindcss(),
    ],
    publicDir: false,
    build: {
        outDir: 'public/dist',
        // generate manifest.json in outDir
        manifest: true,
        rollupOptions: {
            // overwrite default .html entry
            input: '/assets/js/main.js',
        },
    },
    resolve: {
        alias: {
            '@': '/resources/js',
        },
    },
    esbuild: {
        globalName: "studio",
        format: "iife"
    },
    server: {
        // origin: 'http://127.0.0.1:8000',
        proxy: {
            '/foo': 'http://localhost:8000',
        },
    },
});
