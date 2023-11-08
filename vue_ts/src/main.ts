import { InjectionKey, createApp } from "vue";
import App from "./App.vue";
import "@/index.css";
import router from "./router";
import repository from "@/apis/repositoryFactory";
import repositoryFactory from "@/apis/repositoryFactory";

// createApp(App).use(router).mount("#app");
const app = createApp(App).use(router);
app.config.globalProperties.$repository = repository;

//
// `export const ***`は、***.tsでしかできません（***.vueでは無理）
//
export const key: InjectionKey<typeof repositoryFactory> = Symbol("repository");

app.mount("#app");
