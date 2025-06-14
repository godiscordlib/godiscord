<template>
    <div class="h-screen text-blue-100 w-full gap-[2.3rem] bg-zinc-900 flex flex-col justify-center items-center py-4">
        <div class="flex flex-col items-center">
            <NuxtImg src="/godiscord.webp" width="144" height="144" class="transition-all hover:drop-shadow-lg hover:drop-shadow-[#00ADD8] mb-4" />
            <h2 class="text-3xl font-semibold">godiscord – Simple. Powerful. <span class="text-[#00ADD8]">Go</span>.</h2>
            <p class="mt-2">godiscord is a simple and performant Discord library to make Discord bots.</p>


        </div>
      <span class="bg-zinc-950 px-4 py-2 rounded-sm font-mono">$ go get godiscord.foo.ng/lib</span>
      <div class="flex gap-6">
        <div class="bg-[#181818] text-white text-[1.05rem] rounded-sm ring-1 ring-gray-600 shadow-md px-4 py-3">
          <h3>🤏 Small size</h3>
        </div>
        <div class="bg-[#181818] text-white text-[1.05rem] rounded-sm ring-1 ring-gray-600 shadow-md px-4 py-3">
          <h3>⚡ Blazingly fast</h3>
        </div>
        <div class="bg-[#181818] text-white text-[1.05rem] rounded-sm ring-1 ring-gray-600 shadow-md px-4 py-3">
          <h3>🛠 Easy to use</h3>
        </div>
      </div>
      <div class="flex gap-6">
        <NuxtLink to="/docs" class="bg-blue-500 px-4 py-2 rounded-sm text-[1.05rem] ring-1 ring-blue-400">See the documentation</NuxtLink>
        <NuxtLink to="/discord" class="bg-[#181818] px-4 py-2 rounded-sm text-[1.05rem] ring-1 ring-gray-600">Join the Discord</NuxtLink>
      </div>
    </div>
</template>
<style>
body {
    font-family: Inter, sans-serif;
}
</style>