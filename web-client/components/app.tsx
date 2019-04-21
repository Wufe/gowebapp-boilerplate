import * as React from 'react';
import './app.scss';

export const App = () => <div className="centered-landing-content">
	<div className="app__title">Boilerplate</div>
	<div className="app__subtitle">Starting point for my next project</div>
	<div className="app__impressions">
		<div className="impression first">Handlers - Services - Models</div>
		<div className="impression second">Static public folder</div>
		<div className="impression third">Dependency injection</div>
		<div className="impression fourth">Client (Typescript)</div>
		<div className="impression fifth">SQLITE ORM</div>
	</div>
	
</div>;