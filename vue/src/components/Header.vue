<template>
    <header class="bg-white">
        <nav class="mx-auto flex max-w-3xl items-center justify-between p-6 lg:px-8" aria-label="Global">
            <div class="flex lg:flex-1">
                <a href="/" class="-m-1.5 p-1.5 text-3xl">
                    <span class="sr-only">Your Company</span>
                    <svg xmlns="http://www.w3.org/2000/svg" height="1em" viewBox="0 0 512 512">
                        <path d="M226.5 92.9c14.3 42.9-.3 86.2-32.6 96.8s-70.1-15.6-84.4-58.5s.3-86.2 32.6-96.8s70.1 15.6 84.4 58.5zM100.4 198.6c18.9 32.4 14.3 70.1-10.2 84.1s-59.7-.9-78.5-33.3S-2.7 179.3 21.8 165.3s59.7 .9 78.5 33.3zM69.2 401.2C121.6 259.9 214.7 224 256 224s134.4 35.9 186.8 177.2c3.6 9.7 5.2 20.1 5.2 30.5v1.6c0 25.8-20.9 46.7-46.7 46.7c-11.5 0-22.9-1.4-34-4.2l-88-22c-15.3-3.8-31.3-3.8-46.6 0l-88 22c-11.1 2.8-22.5 4.2-34 4.2C84.9 480 64 459.1 64 433.3v-1.6c0-10.4 1.6-20.8 5.2-30.5zM421.8 282.7c-24.5-14-29.1-51.7-10.2-84.1s54-47.3 78.5-33.3s29.1 51.7 10.2 84.1s-54 47.3-78.5 33.3zM310.1 189.7c-32.3-10.6-46.9-53.9-32.6-96.8s52.1-69.1 84.4-58.5s46.9 53.9 32.6 96.8s-52.1 69.1-84.4 58.5z"/>
                    </svg>
                    <!-- <img class="h-8 w-auto" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600" alt="" /> -->
                </a>
            </div>
            <div class="flex lg:hidden">
                <button type="button" class="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700" @click="mobileMenuOpen = true">
                <span class="sr-only">Open main menu</span>
                <Bars3Icon class="h-6 w-6" aria-hidden="true" />
                </button>
            </div>
            <PopoverGroup class="hidden lg:flex lg:gap-x-12">
                <Popover class="relative">
                    <PopoverButton class="flex items-center gap-x-1 text-sm font-semibold leading-6 text-gray-900">
                        Menu
                        <ChevronDownIcon class="h-5 w-5 flex-none text-gray-400" aria-hidden="true" />
                    </PopoverButton>
                    <transition enter-active-class="transition ease-out duration-200" enter-from-class="opacity-0 translate-y-1" enter-to-class="opacity-100 translate-y-0" leave-active-class="transition ease-in duration-150" leave-from-class="opacity-100 translate-y-0" leave-to-class="opacity-0 translate-y-1">
                        <PopoverPanel class="absolute -left-8 top-full z-10 mt-3 w-screen max-w-md overflow-hidden rounded-3xl bg-white shadow-lg ring-1 ring-gray-900/5">
                        <div class="p-4">
                            <div v-for="item in menu" :key="item.name" class="group relative flex items-center gap-x-6 rounded-lg p-4 text-sm leading-6 hover:bg-gray-50">
                            <div class="flex h-11 w-11 flex-none items-center justify-center rounded-lg bg-gray-50 group-hover:bg-white">
                                <component :is="item.icon" class="h-6 w-6 text-gray-600 group-hover:text-indigo-600" aria-hidden="true" />
                            </div>
                            <div class="flex-auto">
                                <a :href="item.href" class="block font-semibold text-gray-900">
                                    {{ item.name }}
                                    <span class="absolute inset-0" />
                                </a>
                                <p class="mt-1 text-gray-600">{{ item.description }}</p>
                            </div>
                            </div>
                        </div>
                        </PopoverPanel>
                    </transition>
                </Popover>
                <Popover class="relative">
                    <PopoverButton class="flex items-center gap-x-1 text-sm font-semibold leading-6 text-gray-900">
                        Add
                        <ChevronDownIcon class="h-5 w-5 flex-none text-gray-400" aria-hidden="true" />
                    </PopoverButton>
                    <transition enter-active-class="transition ease-out duration-200" enter-from-class="opacity-0 translate-y-1" enter-to-class="opacity-100 translate-y-0" leave-active-class="transition ease-in duration-150" leave-from-class="opacity-100 translate-y-0" leave-to-class="opacity-0 translate-y-1">
                        <PopoverPanel class="absolute -left-8 top-full z-10 mt-3 w-screen max-w-md overflow-hidden rounded-3xl bg-white shadow-lg ring-1 ring-gray-900/5">
                        <div class="p-4">
                            <div v-for="item in addMenu" :key="item.name" class="group relative flex items-center gap-x-6 rounded-lg p-4 text-sm leading-6 hover:bg-gray-50">
                                <div class="flex h-11 w-11 flex-none items-center justify-center rounded-lg bg-gray-50 group-hover:bg-white">
                                    <component :is="item.icon" class="h-6 w-6 text-gray-600 group-hover:text-indigo-600" aria-hidden="true" />
                                </div>
                                <div class="flex-auto">
                                    <a :href="item.href" class="block font-semibold text-gray-900">
                                        {{ item.name }}
                                        <span class="absolute inset-0" />
                                    </a>
                                    <p class="mt-1 text-gray-600">{{ item.description }}</p>
                                </div>
                            </div>
                        </div>
                        </PopoverPanel>
                    </transition>
                </Popover>
                <Popover class="relative">
                    <PopoverButton class="flex items-center gap-x-1 text-sm font-semibold leading-6 text-gray-900">
                        <router-link to="/about">About</router-link>
                    </PopoverButton>
                </Popover>
                <!-- <router-link to="/" class="text-sm font-semibold leading-6 text-gray-900">Home</router-link>
                <router-link  outer-link to="/about" class="text-sm font-semibold leading-6 text-gray-900">About</router-link>
                <router-link to="/create-schedule" class="text-sm font-semibold leading-6 text-gray-900">Create Schedule</router-link> -->
            </PopoverGroup>
            <div class="items-center hidden lg:flex lg:flex-1 lg:justify-end">
                <router-link to="/" class="text-sm font-semibold leading-6 text-gray-900">Home</router-link>
                <HomeIcon class="ms-1 h-5 w-5 flex-none text-gray-400" aria-hidden="true" />
            </div>
        </nav>

        <Dialog as="div" class="lg:hidden" @close="mobileMenuOpen = false" :open="mobileMenuOpen">
            <div class="fixed inset-0 z-10">
                <DialogPanel class="fixed inset-y-0 right-0 z-10 w-full overflow-y-auto bg-white px-6 py-6 sm:max-w-sm sm:ring-1 sm:ring-gray-900/10">
                    <div class="flex items-center justify-between">
                        <a href="/" class="-m-1.5 p-1.5 text-3xl">
                            <span class="sr-only">Your Company</span>
                            <svg xmlns="http://www.w3.org/2000/svg" height="1em" viewBox="0 0 512 512">
                                <path d="M226.5 92.9c14.3 42.9-.3 86.2-32.6 96.8s-70.1-15.6-84.4-58.5s.3-86.2 32.6-96.8s70.1 15.6 84.4 58.5zM100.4 198.6c18.9 32.4 14.3 70.1-10.2 84.1s-59.7-.9-78.5-33.3S-2.7 179.3 21.8 165.3s59.7 .9 78.5 33.3zM69.2 401.2C121.6 259.9 214.7 224 256 224s134.4 35.9 186.8 177.2c3.6 9.7 5.2 20.1 5.2 30.5v1.6c0 25.8-20.9 46.7-46.7 46.7c-11.5 0-22.9-1.4-34-4.2l-88-22c-15.3-3.8-31.3-3.8-46.6 0l-88 22c-11.1 2.8-22.5 4.2-34 4.2C84.9 480 64 459.1 64 433.3v-1.6c0-10.4 1.6-20.8 5.2-30.5zM421.8 282.7c-24.5-14-29.1-51.7-10.2-84.1s54-47.3 78.5-33.3s29.1 51.7 10.2 84.1s-54 47.3-78.5 33.3zM310.1 189.7c-32.3-10.6-46.9-53.9-32.6-96.8s52.1-69.1 84.4-58.5s46.9 53.9 32.6 96.8s-52.1 69.1-84.4 58.5z"/>
                            </svg>
                            <!-- <img class="h-8 w-auto" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600" alt="" /> -->
                        </a>
                        <button type="button" class="-m-2.5 rounded-md p-2.5 text-gray-700" @click="mobileMenuOpen = false">
                            <span class="sr-only">Close menu</span>
                            <XMarkIcon class="h-6 w-6" aria-hidden="true" />
                        </button>
                    </div>
                    <div class="mt-6 flow-root">
                        <div class="-my-6 divide-y divide-gray-500/10">
                            <div class="space-y-2 py-6">
                                <Disclosure as="div" class="-mx-3" v-slot="{ open }">
                                    <DisclosureButton class="flex w-full items-center justify-between rounded-lg py-2 pl-3 pr-3.5 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">
                                        Menu
                                        <ChevronDownIcon :class="[open ? 'rotate-180' : '', 'h-5 w-5 flex-none']" aria-hidden="true" />
                                    </DisclosureButton>
                                    <DisclosurePanel class="mt-2 space-y-2">
                                        <DisclosureButton v-for="item in [...menu]" :key="item.name" as="a" :href="item.href" class="block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50">
                                            {{ item.name }}
                                            <p class="mt-1 text-gray-500 text-xs">{{ item.description }}</p>
                                        </DisclosureButton>
                                    </DisclosurePanel>
                                </Disclosure>
                                <Disclosure as="div" class="-mx-3" v-slot="{ open }">
                                    <DisclosureButton class="flex w-full items-center justify-between rounded-lg py-2 pl-3 pr-3.5 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">
                                        Add
                                        <ChevronDownIcon :class="[open ? 'rotate-180' : '', 'h-5 w-5 flex-none']" aria-hidden="true" />
                                    </DisclosureButton>
                                    <DisclosurePanel class="mt-2 space-y-2">
                                        <DisclosureButton v-for="item in [...addMenu]" :key="item.name" as="a" :href="item.href" class="block rounded-lg py-2 pl-6 pr-3 text-sm font-semibold leading-7 text-gray-900 hover:bg-gray-50">
                                            {{ item.name }}
                                            <p class="mt-1 text-gray-500 text-xs">{{ item.description }}</p>
                                        </DisclosureButton>
                                    </DisclosurePanel>
                                </Disclosure>
                                <Disclosure as="div" class="-mx-3">
                                    <DisclosureButton class="flex w-full items-center justify-between rounded-lg py-2 pl-3 pr-3.5 text-base font-semibold leading-7 text-gray-900 hover:bg-gray-50">
                                        About
                                    </DisclosureButton>
                                </Disclosure>
                                <!-- <router-link to="/" class="text-sm font-semibold leading-6 text-gray-900">Home</router-link>
                                <router-link  outer-link to="/about" class="text-sm font-semibold leading-6 text-gray-900">About</router-link>
                                <router-link to="/create-schedule" class="text-sm font-semibold leading-6 text-gray-900">Create Schedule</router-link> -->
                            </div>
                            <div class="py-6 flex items-start">
                                <router-link to="/" class="text-sm font-semibold leading-6 text-gray-900">Home</router-link>
                                <HomeIcon class="ms-1 h-5 w-5 flex-none text-gray-400" aria-hidden="true" />
                            </div>
                        </div>
                    </div>
                </DialogPanel>
            </div>
        </Dialog>
    </header>
</template>

<script setup>
    import { ref } from 'vue'
    import {
        Dialog,
        DialogPanel,
        Disclosure,
        DisclosureButton,
        DisclosurePanel,
        Popover,
        PopoverButton,
        PopoverGroup,
        PopoverPanel,
    } from '@headlessui/vue'
    import {
        Bars3Icon,
        ChartBarIcon,
        PencilIcon,
        CalendarIcon,
        SquaresPlusIcon,
        XMarkIcon,
        HomeIcon,
    } from '@heroicons/vue/24/outline'
    import { ChevronDownIcon } from '@heroicons/vue/20/solid'

    const menu = [
        { name: 'Weight', description: '体重の推移', href: '#', icon: ChartBarIcon },
        { name: 'Memo', description: 'メモ一覧', href: '#', icon: PencilIcon },
        { name: 'Schedule', description: '予定一覧', href: '#', icon: CalendarIcon },
    ]
    const addMenu = [
        { name: 'Weight', description: '体重を記録する', href: '#', icon: ChartBarIcon },
        { name: 'Memo', description: 'メモを追加する', href: '#', icon: PencilIcon },
        { name: 'Schedule', description: '予定を追加する', href: '/create-schedule', icon: SquaresPlusIcon },
    ]
    const mobileMenuOpen = ref(false)
</script>