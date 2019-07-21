// Project Type
type Project = {
	_id: string;
	title: string;
	description: string;
	description_link: string;
	html: string;
	markdown: string;
	doc_type: string;
	thumbnail: string;
	link: string;
	created_at: number;
	is_major: boolean;
	is_public: boolean;
	is_deleted: boolean;
};

export default Project;
