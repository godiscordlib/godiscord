// tailwind.config.js
module.exports = {
    content: [
        './components/**/*.{vue,js,ts}',
        './layouts/**/*.vue',
        './pages/**/*.vue',
        './app.vue',
        './nuxt.config.{js,ts}'
    ],
    theme: {
        extend: {
            colors: {
                primary: '#1e293b', // Gris foncé stylé 🖤
                accent: '#22d3ee', // Cyan 💎
                danger: '#ef4444' // Rouge 🔥
            },
            fontFamily: {
                custom: ['"Poppins"', 'sans-serif']
            }
        }
    }
}
