import React, { useState, useEffect } from 'react';

import SubBoxName from './SubBoxName';
import SubBoxDoc from './SubBoxDoc';
import SubBoxBool from './SubBoxBool';

import Project from '../types/project';

type MyProps = {
	project: Project;
	pContext: any;
	setProject: () => void;
};

const ProjectBoxes = (props: MyProps) => {
	const { project } = props;
	const [ title, setTitle ] = useState(project ? project.title : '');
	const [ description, setDescription ] = useState(project ? project.description : '');
	const [ link, setLink ] = useState(project ? project.link : '');

	const [ isAsync, setIsAsync ] = useState(false); // Is Async

	const [ nameBoxEd, setNameBoxEd ] = useState(false); // Check if Title, Emoji, or Desc. has changed
	// const [ nameBoxEd, setNameBoxEd ] = useState();	// Check if Title, Emoji, or Desc. has changed

	return (
		<div className="ProjectBoxes-c-00">
			<SubBoxName object={project} />
			<SubBoxDoc object={project} />
			<SubBoxBool object={project} />
		</div>
	);
};

export default ProjectBoxes;
