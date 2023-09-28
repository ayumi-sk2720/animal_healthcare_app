<template>
    <footer class="bg-emerald-500 text-white text-xs lg:text-sm mt-6">
        <div class="mx-auto block lg:flex max-w-3xl items-center justify-between p-6 lg:px-8" aria-label="Global">
            <PopoverGroup class="block mb-4 lg:m-0 lg:flex lg:gap-x-6">
                <Popover class="relative">
                    <PopoverButton class="flex items-center gap-x-1 leading-6">
                        <router-link to="/about">About</router-link>
                    </PopoverButton>
                </Popover>
                <Popover class="relative">
                    <PopoverButton class="flex items-center gap-x-1 leading-6">
                        <router-link to="/about">サイトマップ</router-link>
                    </PopoverButton>
                </Popover>
                <Popover class="relative">
                    <PopoverButton class="flex items-center gap-x-1 leading-6">
                        <router-link to="/about">プライバシーポリシー</router-link>
                    </PopoverButton>
                </Popover>
            </PopoverGroup>
            <div class="text-right">
                <p class="copyright">© 2023 Example Inc. All Rights Reserved.</p>
            </div>
        </div>
    </footer>
</template>

<script setup>
    import {
        Popover,
        PopoverButton,
        PopoverGroup,
    } from '@headlessui/vue'
</script>