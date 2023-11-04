import { AxiosPromise } from "axios";
import repository from "./repository";

const resource = "/pet";

export interface Pet {
  name: string;
  age: string;
  sex: string;
  nowWeight: string;
  targetWeight: string;
  birthDay: string;
}

export default class PetRepository {
  // TODO: AxiosPromise<Pet>にならない？
  public getPetSummary(id: number): AxiosPromise<Pet> {
    return repository.get<Pet>(`${resource}/${id}`);
  }
}
