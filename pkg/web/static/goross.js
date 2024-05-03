const {createApp, ref} = Vue;

const app = createApp({
    setup() {
        const message = ref('Hello vue!');
        return {
            message
        };
    }
});

app.use(Quasar);

app.mount('#app');
