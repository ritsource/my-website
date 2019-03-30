// Project Type
type Project = {
	_id: string;
	title: string;
	description: string;
	html: string;
	markdown: string;
	doc_type: string;
	link: string;
	image_url: string;
	is_public: boolean;
	is_deleted: boolean;
};

export default Project;
