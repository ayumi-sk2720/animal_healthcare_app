import { App, InjectionKey } from "vue";
import repositoryFactory from "./apis/repositoryFactory";

export const key: InjectionKey<typeof repositoryFactory> = Symbol("repository");
export const register = (app: App) => app.provide(key, repositoryFactory);
