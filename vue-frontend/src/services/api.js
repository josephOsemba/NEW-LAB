import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:5000/api',
});

export const getHello = async () => {
  const response = await api.get('/hello');
  return response.data;
}
