import React, { useState, useEffect } from 'react';

import JSONBox from '../components/JSONBox';
import ProjectBoxes from '../components/ProjectBox';
import data, { Data } from '../data/data';

import ProjectContext from '../contexts/ProjectContext';

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

	useEffect(() => {
		setProject(pContext.projects.find(({ _id }: any) => _id === props.match.params.projectId));
	}, []);

	useEffect(
		() => {
			console.log(project);
		},
		[ project ]
	);

	return (
		<div className="Page-c-00">
			<div className="Page-Container-00">
				<h1>Project / {props.match.params.projectId}</h1>
				<div style={{ alignItems: 'flex-start' }} className="Flex-Row-Space-Between">
					<JSONBox object={project} />
					<div className="Page-Vertical-Box-Container">
						<ProjectBoxes pContext={pContext} project={project} setProject={setProject} />
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
