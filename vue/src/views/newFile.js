import AjaxRequestWrapper from '@/ajaxWrapper';

export default (await import('vue')).defineComponent({
data() {
return {
schedule: [],
title: '',
date: '',
place: '',
};
},
mounted() {
},
methods: {
onClickPost: function (e) {
console.log(this.title); // トリミング
console.log(this.date); // 2023-10-27T12:13
console.log(this.place); // hogehoge

new AjaxRequestWrapper(
'http://localhost:8080/pet/2/schedule',
{
title: this.title,
date: this.date,
location: this.place,
},
"aaaaa"


);
}
}
});
