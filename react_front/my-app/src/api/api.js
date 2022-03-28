import axios from 'axios'
const API_URI = 'http://localhost:8080'


const api = axios.create({
    baseURL: API_URI,
    headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
        

    }
});

export default api