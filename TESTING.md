# TrackP Testing Guide

## 🧪 Manual Testing Checklist

### Prerequisites
- ✅ Backend server running on port 8080 (demo mode)
- ✅ Frontend server running on port 3000
- ✅ Browser open at http://localhost:3000

### 🏠 Dashboard Features

#### Project Management
- [ ] **Create Project**: Click "Add New Project" and create a test project
- [ ] **Edit Project**: Click edit button and modify project details
- [ ] **Delete Project**: Test project deletion with confirmation
- [ ] **View Tasks**: Click "View Tasks" to navigate to Kanban board

#### 🔍 Search Functionality
- [ ] **Search Projects**: Type in the search box and verify real-time filtering
- [ ] **Search Performance**: Verify debounced input (no lag during typing)
- [ ] **Clear Search**: Test the clear button (×) functionality
- [ ] **Case Insensitive**: Search with different cases (e.g., "TEST", "test", "Test")
- [ ] **Partial Matches**: Search for partial project names and descriptions

#### 📊 Sorting Features
- [ ] **Sort by Name**: Click "Sort by Name" and verify alphabetical order
- [ ] **Sort by Date**: Click "Sort by Date" and verify chronological order (newest first)
- [ ] **Sort Persistence**: Refresh page and verify sort preference is remembered
- [ ] **Sort Indicators**: Verify sort direction indicators are displayed

#### 💾 Data Persistence
- [ ] **LocalStorage**: Create projects, refresh page, verify data persists
- [ ] **Search State**: Perform search, refresh, verify search term is cleared
- [ ] **Sort State**: Set sort preference, refresh, verify sort is maintained

### 🎯 Kanban Board Features

#### Task Management
- [ ] **Create Task**: Add new tasks in different columns
- [ ] **Edit Task**: Modify task title and description
- [ ] **Delete Task**: Remove tasks and verify updates
- [ ] **Task Details**: Verify all task information displays correctly

#### 🖱️ Drag & Drop
- [ ] **Column Movement**: Drag tasks between different columns
- [ ] **Same Column**: Reorder tasks within the same column
- [ ] **Visual Feedback**: Verify drag indicators and drop zones
- [ ] **API Updates**: Confirm changes are saved to backend

#### 📱 Responsive Design
- [ ] **Desktop View**: Test on full-screen browser
- [ ] **Tablet View**: Resize browser to tablet dimensions
- [ ] **Mobile View**: Test mobile-responsive layout
- [ ] **Touch Devices**: If available, test touch interactions

### 🔧 Technical Features

#### Error Handling
- [ ] **API Errors**: Stop backend server and test error messages
- [ ] **Network Issues**: Test offline behavior (if useOnlineStatus is active)
- [ ] **Invalid Data**: Try submitting empty forms or invalid data
- [ ] **Loading States**: Verify loading spinners appear during operations

#### Performance
- [ ] **Search Debouncing**: Type rapidly and verify smooth performance
- [ ] **Large Datasets**: Create multiple projects and test performance
- [ ] **Memory Leaks**: Navigate between views multiple times
- [ ] **Browser Console**: Check for JavaScript errors or warnings

### 🎨 User Experience

#### Visual Design
- [ ] **Layout Consistency**: Verify spacing and alignment
- [ ] **Color Scheme**: Check contrast and readability
- [ ] **Icons and Buttons**: Test all interactive elements
- [ ] **Animations**: Verify smooth transitions and hover effects

#### Accessibility
- [ ] **Keyboard Navigation**: Tab through all interactive elements
- [ ] **Screen Reader**: Test ARIA labels (if screen reader available)
- [ ] **Focus Indicators**: Verify visible focus states
- [ ] **Color Accessibility**: Test with high contrast mode

### 🚀 Advanced Features

#### Custom Hooks Testing
- [ ] **useLocalStorage**: Verify data persistence across sessions
- [ ] **useDebounce**: Test search input debouncing
- [ ] **useAsync**: Check loading states and error handling
- [ ] **useForm**: Test form validation and submission

#### Utility Functions
- [ ] **Date Formatting**: Check project creation dates display correctly
- [ ] **Error Messages**: Verify user-friendly error messages
- [ ] **Data Validation**: Test input validation and sanitization

## 🐛 Common Issues to Watch For

### Frontend Issues
- [ ] **Console Errors**: Check browser console for JavaScript errors
- [ ] **React Warnings**: Look for development warnings in console
- [ ] **TypeScript Errors**: Verify no TypeScript compilation errors
- [ ] **CSS Issues**: Check for layout problems or missing styles

### Backend Issues
- [ ] **API Responses**: Verify correct HTTP status codes
- [ ] **CORS Issues**: Check for cross-origin request problems
- [ ] **Data Consistency**: Ensure database/memory state is consistent
- [ ] **Error Handling**: Test API error responses

### Performance Issues
- [ ] **Slow Loading**: Monitor load times for components
- [ ] **Memory Usage**: Check for memory leaks in long sessions
- [ ] **Network Requests**: Verify efficient API usage
- [ ] **Bundle Size**: Monitor for excessive JavaScript bundle size

## 📊 Test Results Template

### ✅ Passed Tests
```
✅ Project creation and editing
✅ Search functionality with debouncing
✅ Sort by name and date
✅ LocalStorage persistence
✅ Kanban drag & drop
✅ Responsive design
✅ Error handling
```

### ❌ Failed Tests
```
❌ [Test Name]: [Issue Description]
❌ [Test Name]: [Issue Description]
```

### 🔍 Observations
```
💡 [Performance notes]
💡 [User experience insights]
💡 [Potential improvements]
```

## 🎯 Success Criteria

### Must Pass
- All CRUD operations work correctly
- Search and sort functions properly
- Drag & drop saves changes to backend
- No console errors or warnings
- Responsive design works on all screen sizes

### Nice to Have
- Smooth animations and transitions
- Fast search performance
- Persistent user preferences
- Accessible keyboard navigation
- Helpful error messages

## 🔄 Test Automation Prep

Future test automation should cover:
- API endpoint testing
- Component unit tests
- E2E user flow testing
- Performance benchmarking
- Cross-browser compatibility

---

**Test Status**: Ready for manual testing
**Last Updated**: 2025-07-27
**Version**: TrackP v1.1.0
