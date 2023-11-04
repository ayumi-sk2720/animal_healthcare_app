import PetRepository from "./petRepository";

export interface Repositories {
  pet: PetRepository;
}

function getRepositories(): Repositories {
  const pet = new PetRepository();
  const repositories: Repositories = {
    pet,
  };
  return repositories;
}

export default getRepositories();
