import React, { useState, useEffect } from 'react';

import SubBoxName from './SubBoxName';
import SubBoxDoc from './SubBoxDoc';
import SubBoxBool from './SubBoxBool';

import Project from '../types/project';

type MyProps = {
	project: Project;
	pContext: any;
	setProject: () => void;
	saveFunction: (c: Project, u: Project) => void;
};

const ProjectBoxes = (props: MyProps) => {
	const { project, saveFunction } = props;

	return (
		<div className="ProjectBoxes-c-00">
			<SubBoxName object={project} saveFunction={saveFunction} />
			<SubBoxDoc object={project} saveFunction={saveFunction} />
			<SubBoxBool object={project} saveFunction={saveFunction} />
		</div>
	);
};

export default ProjectBoxes;
