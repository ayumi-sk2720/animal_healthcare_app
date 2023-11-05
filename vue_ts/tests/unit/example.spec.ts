import { shallowMount } from "@vue/test-utils";
import HelloWorld from "@/components/HelloWorld.vue";

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

describe("sample test", () => {
  test("two plus two is four", () => {
    expect(1 + 2).toBe(4);
  });
});
