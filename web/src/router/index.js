import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
	{
		path: '/',
		redirect: () => {
			return { path: '/login' }
		}
	},
	{
		path: '/login',
		component: () => import(/* webpackChunkName: "about" */ '@/views/login.vue')
	},
	{
		path: '/index',
		redirect: '/sessions',
		component: () => import(/* webpackChunkName: "system" */ '@/views/index.vue'),
		children: [
			{
				path: '/sessions',
				component: () => import(/* webpackChunkName: "system" */ '@/views/sessions.vue'),
				children: [
					{
						path: '/sessions/:id', 
						component: () => import(/* webpackChunkName: "system" */ '@/views/session/view')
					},
				]
			},
			{
				path: '/users',
				component: () => import(/* webpackChunkName: "system" */ '@/views/users.vue'),
				children: [
					{
						path: '/users/:id',
						component: () => import(/* webpackChunkName: "system" */ '@/views/user/view-user')
					},
					{
						path: '/groups/:id',
						component: () => import(/* webpackChunkName: "system" */ '@/views/user/view-group')
					}
				]
			},
			{
				path: '/requests',
				component: () => import(/* webpackChunkName: "system" */ '@/views/requests.vue'),
				children: [
					{
						path: '/requests/:id',
						component: () => import(/* webpackChunkName: "system" */ '@/views/request/view')
					}
				]
			},
			{
				path: '/moments',
				component: () => import(/* webpackChunkName: "system" */ '@/views/moments.vue'),
			},
			{
				path: '/setting',
				component: () => import(/* webpackChunkName: "system" */ '@/views/setting.vue'),
			},
		]
	}
]

const router = createRouter({
	history: createWebHashHistory(),
	routes
})
export default router;
