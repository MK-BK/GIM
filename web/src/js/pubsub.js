export default {
    ws: null,
    subscribers: {},

    connect() {
        let that = this
        let ws = new WebSocket("ws://localhost:8080/api/event?token=" + sessionStorage.getItem('token'))

        ws.onmessage = function(msg) {
            let data = JSON.parse(msg.data);
            let topic = data.Topic;

            if (that.subscribers[topic] && that.subscribers[topic].length) {
                that.subscribers[topic].forEach(fn => {
                    fn(data.Body)
                });
            }
        }
        return ws;
    },

    sub(topic, fn) {
        let ws
        if (!ws) {
            ws = this.connect();
        }
    
        let fns = this.subscribers[topic];
        if (fns && fns.length) {
            fns.push(fn);
            return;
        }

        this.subscribers[topic] = [];
        this.subscribers[topic].push(fn);
    },
  
    unsub: function(topic, fn) {
        let fns = this.subscribers[topic];
        if (fns) {
            var index = fns.indexOf(fn);
            if (index !== -1) {
                fns.splice(index, 1);
                if (!fns.length) {
                    delete this.subscribers[topic];
                }
            }
        }
    },
}