import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { browserHistory } from 'react-router';
import { selectTab } from '../../../ducks/menu';
import TabMenu from '../components';

class TabMenuContainer extends Component {
  selectTab = tabName => {
    if (tabName === '/game/orders') {
      console.log('Toggle display orders');
      return
    }
    this.props.actions.selectTab(tabName);
    browserHistory.replace(tabName);
  }

  showTabs = path => !['/login', '/signup'].includes(path);

  render() {
    const { selectedTab, path } = this.props.state;
    const component = (this.showTabs(path))
      ? <TabMenu selectedTab={selectedTab} selectTab={this.selectTab} />
      : <div />
    return component
  }
}

const mapStateToProps = state => ({
  state: {
    selectedTab: state.menu.selectedTab,
    path: state.routing.locationBeforeTransitions.pathname
  }
});

const mapDispatchToActions = dispatch => ({
  actions: bindActionCreators({ selectTab }, dispatch)
});

export default connect(
  state => mapStateToProps(state),
  dispatch => mapDispatchToActions(dispatch)
)(TabMenuContainer);
