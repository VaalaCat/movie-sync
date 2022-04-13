import axios from 'axios';
import QS from 'qs'

axios.defaults.baseURL = '/api';

axios.interceptors.request.use(
	request => {
		request.headers.authorization = "Bearer " + localStorage["token"];
		return request;
	},
	err => {
		return Promise.reject(err);
	}
);

axios.interceptors.response.use(
	response => {
		if (response.status === 200) {
			if (typeof (response.data.token) != "undefined") {
				localStorage["token"] = response.data.token;
			}
			return Promise.resolve(response);
		} else {
			return Promise.reject(response);
		}
	},
	error => {
		if (error.response.status) {
			return Promise.reject(error.response);
		}
	}
);

/**
 * get方法，对应get请求
 * @param {String} url [请求的url地址]
 * @param {Object} params [请求时携带的参数]
 */
export function get(url, params) {
	return new Promise((resolve, reject) => {
		axios.get(url, {
			params: params
		}).then(res => {
			resolve(res);
		}).catch(err => {
			reject(err)
		})
	});
}

/** 
 * post方法，对应post请求 
 * @param {String} url [请求的url地址] 
 * @param {Object} params [请求时携带的参数] 
 */
export function post(url, params) {
	return new Promise((resolve, reject) => {
		axios.post(url, QS.stringify(params))
			.then(res => {
				resolve(res);
			})
			.catch(err => {
				reject(err)
			})
	});
}
