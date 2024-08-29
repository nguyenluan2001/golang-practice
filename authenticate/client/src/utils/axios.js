import axios from "axios"

axios.defaults.withCredentials = true;

const axiosInstance = axios.create({
    baseURL:'http://localhost:8081',
    // withCredentials: true, // default
    // crossdomain: true,
    headers:{
            'Content-Type' : 'application/x-www-form-urlencoded; charset=UTF-8',
            // "Access-Control-Allow-Origin":"*"
    }
})
export {axiosInstance}