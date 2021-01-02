<template>
    <div>
        <ul>
            <li v-for="task in tasks"
            v-bind:key="task.title">
                {{ task.title }}
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    data() {
        return {
            title: "Hello",
            user: {
                firstName: "Pedro",
                lastName: "Rolando",
            },
            myHtml: '<h3>hola vue</h3>',
            showme: true,
            tasks: [
                {title: "eat"},
                {title: "dinner"},
                {title: "play"}
            ]
        }
    }
}
</script>

<!-- styles section -->
<style scoped>

</style>
