import React, { useState } from 'react';
import Project from '../types/project';

const ProjectContext = React.createContext({
	projects: [],
	readProjects: (a: Array<Project>) => {},
	addProject: (p: Project) => {},
	updateProject: (p: Project) => {},
	deleteProject: (id: string) => {}
});

export const ProjectProvider = (props: any) => {
	const [ projects, setProjects ] = useState([]);

	const readProjects = (allProjects: any) => {
		setProjects(allProjects);
	};

	const addProject = (newProject: any) => {
		setProjects(projects.concat(newProject));
	};

	const updateProject = (newProject: any) => {
		setProjects(projects.filter(({ _id }) => _id !== newProject._id).concat(newProject));
	};

	const deleteProject = (projectId: string) => {
		setProjects([ ...projects.filter(({ _id }) => _id !== projectId) ]);
	};

	return (
		<ProjectContext.Provider
			value={{
				projects,
				readProjects,
				addProject,
				updateProject,
				deleteProject
			}}
		>
			{props.children}
		</ProjectContext.Provider>
	);
};

export default ProjectContext;
