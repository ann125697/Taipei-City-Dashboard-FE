import { defineConfig, splitVendorChunkPlugin } from "vite";
import vue from "@vitejs/plugin-vue";
import viteCompression from "vite-plugin-compression";

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [vue(), viteCompression(), splitVendorChunkPlugin()],
	build: {
		rollupOptions: {
			output: {
				manualChunks(id) {
					if (id.includes("node_modules")) {
						return id
							.toString()
							.split("node_modules/")[1]
							.split("/")[0]
							.toString();
					}
				},
			},
		},
		chunkSizeWarningLimit: 1600,
	},
	base: "/",
	server: {
		host: "0.0.0.0",
		port: process.env.VITE_FE_PORT || 8080,
		proxy: {
			"/api/dev": {
				target: process.env.VITE_BE_URL || "http://dashboard-be:8888",
				changeOrigin: true,
				rewrite: (path) => path.replace("/dev", "/v1"),
			},
			"/geo_server": {
				target: "https://geoserver.tuic.gov.taipei/geoserver/",
				changeOrigin: true,
				rewrite: (path) => path.replace(/^\/geo_server/, ""),
			},
		},
	},
});
