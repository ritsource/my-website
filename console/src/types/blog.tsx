// Blog Type
type Blog = {
	_id: string;
	title: string;
	description: string;
	html: string;
	markdown: string;
	doc_type: string;
	thumbnail: string;
	is_public: boolean;
	is_deleted: boolean;
};

export default Blog;
