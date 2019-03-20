export type Data = {
	_id: string;
	title: string;
	type: string;
	emoji: string;
	description: string;
	html: string;
	markdown: string;
	link: string;
	imageUrl: string;
};

const data: Array<Data> = [
	{
		_id: '1',
		title: 'My Life til 19',
		type: 'blog',
		emoji: '🍏',
		description: 'description',
		html: 'http://content.ritwiksaha.com/xyz',
		markdown: 'http://content.ritwiksaha.com/xyz',
		link: 'http://content.ritwiksaha.com/xyz',
		imageUrl: 'http://content.ritwiksaha.com/xyz'
	},
	{
		_id: '2',
		title: 'Scheduler',
		type: 'project',
		emoji: '🍌',
		description: 'description',
		html: 'http://content.ritwiksaha.com/xyz',
		markdown: 'http://content.ritwiksaha.com/xyz',
		link: 'http://content.ritwiksaha.com/xyz',
		imageUrl: 'http://content.ritwiksaha.com/xyz'
	},
	{
		_id: '3',
		title: 'Raspi',
		type: 'project',
		emoji: '🍓',
		description: 'description',
		html: 'http://content.ritwiksaha.com/xyz',
		markdown: 'http://content.ritwiksaha.com/xyz',
		link: 'http://content.ritwiksaha.com/xyz',
		imageUrl: 'http://content.ritwiksaha.com/xyz'
	}
];

export default data;
