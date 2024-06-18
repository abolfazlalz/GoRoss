const {createApp, ref} = Vue;


const app = createApp({
    setup() {
        const socket = io('/');
        console.log(socket);

        socket.on("connect", () => {
            console.log('socket connected:', socket.id);
        });

        socket.on("error", () => {
            console.log('socket error:', socket.id);
        });

        const handleClick = (id) => {
            socket.emit('action', id);
        }

        return {
            handleClick
        };
    }
});

app.use(Quasar);

app.mount('#app');
