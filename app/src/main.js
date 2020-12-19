import Vue from 'vue'
import App from './App.vue'
import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
import './assets/css/app.css'
import '@fortawesome/fontawesome-free/css/all.css'
import '@fortawesome/fontawesome-free/js/all.js'

Vue.config.productionTip = false

Vue.prototype.$postData = async (url = '', method = '',  data = {}) => {
    const response = await fetch(`http://localhost:8200/api/${url}`, {
        method: method,
        data: data,
        mode: 'cors',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    });

    return response.json()
}

new Vue({
  render: h => h(App),
}).$mount('#app')
