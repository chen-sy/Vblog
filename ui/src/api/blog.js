import client from './client';

const BLIG_LIST = (data) => client({
    url: '/blog/list',
    method: 'get',
    data: data
});

export { BLIG_LIST };
