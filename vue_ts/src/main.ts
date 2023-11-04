import { createApp } from "vue";
import App from "./App.vue";
import "@/index.css";
import router from "./router";
import repository from "@/apis/repositoryFactory";

// createApp(App).use(router).mount("#app");
const app = createApp(App).use(router);
app.config.globalProperties.$repository = repository;

app.mount("#app");
