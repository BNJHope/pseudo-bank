import http from 'k6/http';

import { sleep } from 'k6';


export default function () {
    const baseUrl = __ENV.APP_URL;
    http.get(`${baseUrl}/transactions`);
    sleep(1);
}
