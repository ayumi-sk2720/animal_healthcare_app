import { shallowMount } from "@vue/test-utils";
import HelloWorld from "@/components/HelloWorld.vue";
import axios from "axios";
import PetRepository from "@/apis/petRepository";

jest.mock("axios");

const BASE_URL = "http://localhost:8080/";
/***
 * コンポーネントの表示に関する単純なテスト
 */
describe("HelloWorld.vue", () => {
  it("renders props.msg when passed", () => {
    const msg = "new message";
    const wrapper = shallowMount(HelloWorld, {
      props: { msg },
    });
    expect(wrapper.text()).toMatch(msg);
  });
});

/***
 * 任意のメソッドをテストするサンプル
 */
describe("sample test", () => {
  test("two plus two is four", () => {
    expect(1 + 2).toBe(4);
  });
});

/***
 * APIをテストするサンプル
 * [参考] https://vhudyma-blog.eu/3-ways-to-mock-axios-in-jest/
 */
describe("getPetSummary", () => {
  it("API成功テストケース", async () => {
    const petSummary = {
      pet: {
        id: "1",
        name: "なっちゃん",
        age: "",
        sex: "1",
        now_weight: "4000",
        target_weight: "3500",
        birthday: "",
        created_at: "2023-10-30T16:31:33+09:00",
        updated_at: "2023-10-30T16:31:33+09:00",
      },
      dosage_schedules: {
        today: {
          id: "",
          title: "hogehoge",
          date: "2023-11-04T08:56:09.273600841+09:00",
        },
        next: {
          id: "",
          title: "fugafuga",
          date: "2023-11-04T08:56:09.273600882+09:00",
        }
      },
      physical_condition: {
        food: 3,
        sweat: 2,
        toilet: 1,
      },
      memo: {
        id: "",
        title: "今月に入って飲む水の量が\n増えた気がする\n次に通院した時に先生に相談する",
        date: "2023-11-04T08:56:09.273600924+09:00",
      },
      schedules: [
        {
          id: 0,
          pet_id: "1",
          title: "トリミング",
          date: "2023-11-04T08:56:09.273601091+09:00",
          location: "ペテモ立川店",
          created_at: "0001-01-01T00:00:00Z",
          updated_at: "0001-01-01T00:00:00Z",
        },
        {
          id: 0,
          pet_id: "1",
          title: "通院",
          date: "2023-11-04T08:56:09.273601132+09:00",
          location: "ホゲホゲ病院",
          created_at: "0001-01-01T00:00:00Z",
          updated_at: "0001-01-01T00:00:00Z",
        }
      ]
    };
    (axios.get as any).mockResolvedValueOnce(petSummary);
    // when
    const petRepository = new PetRepository();
    const result = petRepository.getPetSummary(1);

    expect(axios.get).toHaveBeenCalledWith(`${BASE_URL}/pet/1`);
    expect(result).toEqual(petSummary);
  });
});
