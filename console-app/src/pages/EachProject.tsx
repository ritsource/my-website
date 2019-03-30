import React, { useState, useEffect } from 'react';
import api from '../api';

import JSONBox from '../components/JSONBox';
import ProjectBoxes from '../components/ProjectBoxes';

import ProjectContext from '../contexts/ProjectContext';

import Project from '../types/project';

type MyProps = {
	history: any;
	match: any;
	pContext: any;
};

const EachProjectPage = (props: MyProps) => {
	const { pContext } = props;

	const [ project, setProject ]: any = useState(
		pContext.projects.find(({ _id }: any) => _id === props.match.params.projectId)
	);

	// useEffect(() => {
	// 	setProject(pContext.projects.find(({ _id }: any) => _id === props.match.params.projectId));
	// }, []);

	useEffect(
		() => {
			setProject(pContext.projects.find(({ _id }: any) => _id === props.match.params.projectId));
		},
		[ pContext.projects ]
	);

	// Edits Project, current => current project object, updates => updated propserties
	const editProject = async (current: Project, updates: Project) => {
		const projectId = current._id;
		console.log('projectId', projectId);
		console.log('current', current);
		console.log('updates', updates);

		delete current._id;
		delete updates._id;

		try {
			const response = await api.put(`/admin/project/edit/${projectId}`, {
				...current,
				...updates
			});
			pContext.updateProject(response.data);
		} catch (error) {
			// return error;
			throw error.message;
		}
	};

	return (
		<div className="Page-c-00">
			<div className="Page-Container-00">
				<div style={{ alignItems: 'flex-start' }} className="Flex-Row-Space-Between">
					<JSONBox object={project} saveFunction={editProject} />
					<div className="Page-Vertical-Box-Container">
						<ProjectBoxes
							pContext={pContext}
							project={project}
							setProject={setProject}
							saveFunction={editProject}
						/>
					</div>
				</div>
			</div>
		</div>
	);
};

export default (props: any) => (
	<ProjectContext.Consumer>
		{(pContext) => <EachProjectPage pContext={pContext} {...props} />}
	</ProjectContext.Consumer>
);
