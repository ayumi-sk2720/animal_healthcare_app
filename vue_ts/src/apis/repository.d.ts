import * as repositoryFactory from "@/apis/repositoryFactory";

declare module "@vue/runtime-core" {
  export interface ComponentCustomProperties {
    $repository: repositoryFactory.Repositories;
  }
}
