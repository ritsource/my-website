// Blog Type
type Blog = {
	_id: string;
	title: string;
	description: string;
	description_link: string;
	author: string;
	formatted_date: string;
	html: string;
	markdown: string;
	doc_type: string;
	thumbnail: string;
	created_at: number;
	is_technical: boolean;
	is_public: boolean;
	is_deleted: boolean;
	is_series: boolean;
	sub_blogs: Array<any>
};

export default Blog;
