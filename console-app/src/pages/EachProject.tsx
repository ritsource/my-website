import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux';

import { createProject, readProjects, editProject } from '../actions/project_actions';
import JSONBox from '../components/JSONBox';
import ProjectBoxes from '../components/ProjectBoxes';
import LoadingPage from './Loading';

import Project from '../types/project';

type MyProps = {
	history: any;
	match: any;
	projects: Array<Project>;
	editProject: (x: string, y: any) => void;
};

const EachProjectPage = (props: MyProps) => {
	const { projects } = props;

	const [ project, setProject ]: any = useState(
		projects.find(({ _id }: any) => _id === props.match.params.projectId)
	);

	useEffect(
		() => {
			setProject(projects.find(({ _id }: any) => _id === props.match.params.projectId));
		},
		[ projects ]
	);

	return (
		<div className="Page-c-00">
			<div className="Page-Container-00">
				<div style={{ alignItems: 'flex-start' }} className="Flex-Row-Space-Between">
					{project ? (
						<React.Fragment>
							{/* In the function below c = current-project-data; u = current-blog-data */}
							<JSONBox
								object={project}
								saveFunction={(c, u) => props.editProject(c._id, { ...c, ...u })}
							/>
							<div className="Page-Vertical-Box-Container">
								<ProjectBoxes
									project={project}
									setProject={setProject}
									saveFunction={(c, u) => props.editProject(c._id, { ...c, ...u })}
								/>
							</div>
						</React.Fragment>
					) : (
						<LoadingPage />
					)}
				</div>
			</div>
		</div>
	);
};

const mapStateToProps = ({ projects }: { projects: Array<Project> }) => ({
	projects: projects || []
});

const mapDispatchToProps = (dispatch: (x: any) => void) => ({
	readProjects: () => dispatch(readProjects()),
	editProject: (x: string, y: any) => dispatch(editProject(x, y))
});

export default connect(mapStateToProps, mapDispatchToProps)(EachProjectPage);
