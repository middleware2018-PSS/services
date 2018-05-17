// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Accounts", testAccounts)
	t.Run("Admins", testAdmins)
	t.Run("Appointments", testAppointments)
	t.Run("Classes", testClasses)
	t.Run("Grades", testGrades)
	t.Run("Notifications", testNotifications)
	t.Run("Parents", testParents)
	t.Run("Payments", testPayments)
	t.Run("Students", testStudents)
	t.Run("Subjects", testSubjects)
	t.Run("Teachers", testTeachers)
	t.Run("Teaches", testTeaches)
	t.Run("Timetables", testTimetables)
}

func TestDelete(t *testing.T) {
	t.Run("Accounts", testAccountsDelete)
	t.Run("Admins", testAdminsDelete)
	t.Run("Appointments", testAppointmentsDelete)
	t.Run("Classes", testClassesDelete)
	t.Run("Grades", testGradesDelete)
	t.Run("Notifications", testNotificationsDelete)
	t.Run("Parents", testParentsDelete)
	t.Run("Payments", testPaymentsDelete)
	t.Run("Students", testStudentsDelete)
	t.Run("Subjects", testSubjectsDelete)
	t.Run("Teachers", testTeachersDelete)
	t.Run("Teaches", testTeachesDelete)
	t.Run("Timetables", testTimetablesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsQueryDeleteAll)
	t.Run("Admins", testAdminsQueryDeleteAll)
	t.Run("Appointments", testAppointmentsQueryDeleteAll)
	t.Run("Classes", testClassesQueryDeleteAll)
	t.Run("Grades", testGradesQueryDeleteAll)
	t.Run("Notifications", testNotificationsQueryDeleteAll)
	t.Run("Parents", testParentsQueryDeleteAll)
	t.Run("Payments", testPaymentsQueryDeleteAll)
	t.Run("Students", testStudentsQueryDeleteAll)
	t.Run("Subjects", testSubjectsQueryDeleteAll)
	t.Run("Teachers", testTeachersQueryDeleteAll)
	t.Run("Teaches", testTeachesQueryDeleteAll)
	t.Run("Timetables", testTimetablesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceDeleteAll)
	t.Run("Admins", testAdminsSliceDeleteAll)
	t.Run("Appointments", testAppointmentsSliceDeleteAll)
	t.Run("Classes", testClassesSliceDeleteAll)
	t.Run("Grades", testGradesSliceDeleteAll)
	t.Run("Notifications", testNotificationsSliceDeleteAll)
	t.Run("Parents", testParentsSliceDeleteAll)
	t.Run("Payments", testPaymentsSliceDeleteAll)
	t.Run("Students", testStudentsSliceDeleteAll)
	t.Run("Subjects", testSubjectsSliceDeleteAll)
	t.Run("Teachers", testTeachersSliceDeleteAll)
	t.Run("Teaches", testTeachesSliceDeleteAll)
	t.Run("Timetables", testTimetablesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Accounts", testAccountsExists)
	t.Run("Admins", testAdminsExists)
	t.Run("Appointments", testAppointmentsExists)
	t.Run("Classes", testClassesExists)
	t.Run("Grades", testGradesExists)
	t.Run("Notifications", testNotificationsExists)
	t.Run("Parents", testParentsExists)
	t.Run("Payments", testPaymentsExists)
	t.Run("Students", testStudentsExists)
	t.Run("Subjects", testSubjectsExists)
	t.Run("Teachers", testTeachersExists)
	t.Run("Teaches", testTeachesExists)
	t.Run("Timetables", testTimetablesExists)
}

func TestFind(t *testing.T) {
	t.Run("Accounts", testAccountsFind)
	t.Run("Admins", testAdminsFind)
	t.Run("Appointments", testAppointmentsFind)
	t.Run("Classes", testClassesFind)
	t.Run("Grades", testGradesFind)
	t.Run("Notifications", testNotificationsFind)
	t.Run("Parents", testParentsFind)
	t.Run("Payments", testPaymentsFind)
	t.Run("Students", testStudentsFind)
	t.Run("Subjects", testSubjectsFind)
	t.Run("Teachers", testTeachersFind)
	t.Run("Teaches", testTeachesFind)
	t.Run("Timetables", testTimetablesFind)
}

func TestBind(t *testing.T) {
	t.Run("Accounts", testAccountsBind)
	t.Run("Admins", testAdminsBind)
	t.Run("Appointments", testAppointmentsBind)
	t.Run("Classes", testClassesBind)
	t.Run("Grades", testGradesBind)
	t.Run("Notifications", testNotificationsBind)
	t.Run("Parents", testParentsBind)
	t.Run("Payments", testPaymentsBind)
	t.Run("Students", testStudentsBind)
	t.Run("Subjects", testSubjectsBind)
	t.Run("Teachers", testTeachersBind)
	t.Run("Teaches", testTeachesBind)
	t.Run("Timetables", testTimetablesBind)
}

func TestOne(t *testing.T) {
	t.Run("Accounts", testAccountsOne)
	t.Run("Admins", testAdminsOne)
	t.Run("Appointments", testAppointmentsOne)
	t.Run("Classes", testClassesOne)
	t.Run("Grades", testGradesOne)
	t.Run("Notifications", testNotificationsOne)
	t.Run("Parents", testParentsOne)
	t.Run("Payments", testPaymentsOne)
	t.Run("Students", testStudentsOne)
	t.Run("Subjects", testSubjectsOne)
	t.Run("Teachers", testTeachersOne)
	t.Run("Teaches", testTeachesOne)
	t.Run("Timetables", testTimetablesOne)
}

func TestAll(t *testing.T) {
	t.Run("Accounts", testAccountsAll)
	t.Run("Admins", testAdminsAll)
	t.Run("Appointments", testAppointmentsAll)
	t.Run("Classes", testClassesAll)
	t.Run("Grades", testGradesAll)
	t.Run("Notifications", testNotificationsAll)
	t.Run("Parents", testParentsAll)
	t.Run("Payments", testPaymentsAll)
	t.Run("Students", testStudentsAll)
	t.Run("Subjects", testSubjectsAll)
	t.Run("Teachers", testTeachersAll)
	t.Run("Teaches", testTeachesAll)
	t.Run("Timetables", testTimetablesAll)
}

func TestCount(t *testing.T) {
	t.Run("Accounts", testAccountsCount)
	t.Run("Admins", testAdminsCount)
	t.Run("Appointments", testAppointmentsCount)
	t.Run("Classes", testClassesCount)
	t.Run("Grades", testGradesCount)
	t.Run("Notifications", testNotificationsCount)
	t.Run("Parents", testParentsCount)
	t.Run("Payments", testPaymentsCount)
	t.Run("Students", testStudentsCount)
	t.Run("Subjects", testSubjectsCount)
	t.Run("Teachers", testTeachersCount)
	t.Run("Teaches", testTeachesCount)
	t.Run("Timetables", testTimetablesCount)
}

func TestHooks(t *testing.T) {
	t.Run("Accounts", testAccountsHooks)
	t.Run("Admins", testAdminsHooks)
	t.Run("Appointments", testAppointmentsHooks)
	t.Run("Classes", testClassesHooks)
	t.Run("Grades", testGradesHooks)
	t.Run("Notifications", testNotificationsHooks)
	t.Run("Parents", testParentsHooks)
	t.Run("Payments", testPaymentsHooks)
	t.Run("Students", testStudentsHooks)
	t.Run("Subjects", testSubjectsHooks)
	t.Run("Teachers", testTeachersHooks)
	t.Run("Teaches", testTeachesHooks)
	t.Run("Timetables", testTimetablesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Accounts", testAccountsInsert)
	t.Run("Accounts", testAccountsInsertWhitelist)
	t.Run("Admins", testAdminsInsert)
	t.Run("Admins", testAdminsInsertWhitelist)
	t.Run("Appointments", testAppointmentsInsert)
	t.Run("Appointments", testAppointmentsInsertWhitelist)
	t.Run("Classes", testClassesInsert)
	t.Run("Classes", testClassesInsertWhitelist)
	t.Run("Grades", testGradesInsert)
	t.Run("Grades", testGradesInsertWhitelist)
	t.Run("Notifications", testNotificationsInsert)
	t.Run("Notifications", testNotificationsInsertWhitelist)
	t.Run("Parents", testParentsInsert)
	t.Run("Parents", testParentsInsertWhitelist)
	t.Run("Payments", testPaymentsInsert)
	t.Run("Payments", testPaymentsInsertWhitelist)
	t.Run("Students", testStudentsInsert)
	t.Run("Students", testStudentsInsertWhitelist)
	t.Run("Subjects", testSubjectsInsert)
	t.Run("Subjects", testSubjectsInsertWhitelist)
	t.Run("Teachers", testTeachersInsert)
	t.Run("Teachers", testTeachersInsertWhitelist)
	t.Run("Teaches", testTeachesInsert)
	t.Run("Teaches", testTeachesInsertWhitelist)
	t.Run("Timetables", testTimetablesInsert)
	t.Run("Timetables", testTimetablesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("AppointmentToStudentUsingStudent", testAppointmentToOneStudentUsingStudent)
	t.Run("AppointmentToTeacherUsingTeacher", testAppointmentToOneTeacherUsingTeacher)
	t.Run("GradeToTeacherUsingTeacher", testGradeToOneTeacherUsingTeacher)
	t.Run("GradeToStudentUsingStudent", testGradeToOneStudentUsingStudent)
	t.Run("PaymentToStudentUsingStudent", testPaymentToOneStudentUsingStudent)
	t.Run("TeachToSubjectUsingSubject", testTeachToOneSubjectUsingSubject)
	t.Run("TeachToTeacherUsingTeacher", testTeachToOneTeacherUsingTeacher)
	t.Run("TeachToClassUsingClass", testTeachToOneClassUsingClass)
	t.Run("TimetableToClassUsingClass", testTimetableToOneClassUsingClass)
	t.Run("TimetableToSubjectUsingSubject", testTimetableToOneSubjectUsingSubject)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ClassToStudents", testClassToManyStudents)
	t.Run("ClassToTeaches", testClassToManyTeaches)
	t.Run("ClassToTimetables", testClassToManyTimetables)
	t.Run("ParentToStudents", testParentToManyStudents)
	t.Run("StudentToAppointments", testStudentToManyAppointments)
	t.Run("StudentToClasses", testStudentToManyClasses)
	t.Run("StudentToGrades", testStudentToManyGrades)
	t.Run("StudentToParents", testStudentToManyParents)
	t.Run("StudentToPayments", testStudentToManyPayments)
	t.Run("SubjectToTeaches", testSubjectToManyTeaches)
	t.Run("SubjectToTimetables", testSubjectToManyTimetables)
	t.Run("TeacherToAppointments", testTeacherToManyAppointments)
	t.Run("TeacherToGrades", testTeacherToManyGrades)
	t.Run("TeacherToTeaches", testTeacherToManyTeaches)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AppointmentToStudentUsingStudent", testAppointmentToOneSetOpStudentUsingStudent)
	t.Run("AppointmentToTeacherUsingTeacher", testAppointmentToOneSetOpTeacherUsingTeacher)
	t.Run("GradeToTeacherUsingTeacher", testGradeToOneSetOpTeacherUsingTeacher)
	t.Run("GradeToStudentUsingStudent", testGradeToOneSetOpStudentUsingStudent)
	t.Run("PaymentToStudentUsingStudent", testPaymentToOneSetOpStudentUsingStudent)
	t.Run("TeachToSubjectUsingSubject", testTeachToOneSetOpSubjectUsingSubject)
	t.Run("TeachToTeacherUsingTeacher", testTeachToOneSetOpTeacherUsingTeacher)
	t.Run("TeachToClassUsingClass", testTeachToOneSetOpClassUsingClass)
	t.Run("TimetableToClassUsingClass", testTimetableToOneSetOpClassUsingClass)
	t.Run("TimetableToSubjectUsingSubject", testTimetableToOneSetOpSubjectUsingSubject)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("GradeToStudentUsingStudent", testGradeToOneRemoveOpStudentUsingStudent)
	t.Run("PaymentToStudentUsingStudent", testPaymentToOneRemoveOpStudentUsingStudent)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("ClassToStudents", testClassToManyAddOpStudents)
	t.Run("ClassToTeaches", testClassToManyAddOpTeaches)
	t.Run("ClassToTimetables", testClassToManyAddOpTimetables)
	t.Run("ParentToStudents", testParentToManyAddOpStudents)
	t.Run("StudentToAppointments", testStudentToManyAddOpAppointments)
	t.Run("StudentToClasses", testStudentToManyAddOpClasses)
	t.Run("StudentToGrades", testStudentToManyAddOpGrades)
	t.Run("StudentToParents", testStudentToManyAddOpParents)
	t.Run("StudentToPayments", testStudentToManyAddOpPayments)
	t.Run("SubjectToTeaches", testSubjectToManyAddOpTeaches)
	t.Run("SubjectToTimetables", testSubjectToManyAddOpTimetables)
	t.Run("TeacherToAppointments", testTeacherToManyAddOpAppointments)
	t.Run("TeacherToGrades", testTeacherToManyAddOpGrades)
	t.Run("TeacherToTeaches", testTeacherToManyAddOpTeaches)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("ClassToStudents", testClassToManySetOpStudents)
	t.Run("ParentToStudents", testParentToManySetOpStudents)
	t.Run("StudentToClasses", testStudentToManySetOpClasses)
	t.Run("StudentToGrades", testStudentToManySetOpGrades)
	t.Run("StudentToParents", testStudentToManySetOpParents)
	t.Run("StudentToPayments", testStudentToManySetOpPayments)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("ClassToStudents", testClassToManyRemoveOpStudents)
	t.Run("ParentToStudents", testParentToManyRemoveOpStudents)
	t.Run("StudentToClasses", testStudentToManyRemoveOpClasses)
	t.Run("StudentToGrades", testStudentToManyRemoveOpGrades)
	t.Run("StudentToParents", testStudentToManyRemoveOpParents)
	t.Run("StudentToPayments", testStudentToManyRemoveOpPayments)
}

func TestReload(t *testing.T) {
	t.Run("Accounts", testAccountsReload)
	t.Run("Admins", testAdminsReload)
	t.Run("Appointments", testAppointmentsReload)
	t.Run("Classes", testClassesReload)
	t.Run("Grades", testGradesReload)
	t.Run("Notifications", testNotificationsReload)
	t.Run("Parents", testParentsReload)
	t.Run("Payments", testPaymentsReload)
	t.Run("Students", testStudentsReload)
	t.Run("Subjects", testSubjectsReload)
	t.Run("Teachers", testTeachersReload)
	t.Run("Teaches", testTeachesReload)
	t.Run("Timetables", testTimetablesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Accounts", testAccountsReloadAll)
	t.Run("Admins", testAdminsReloadAll)
	t.Run("Appointments", testAppointmentsReloadAll)
	t.Run("Classes", testClassesReloadAll)
	t.Run("Grades", testGradesReloadAll)
	t.Run("Notifications", testNotificationsReloadAll)
	t.Run("Parents", testParentsReloadAll)
	t.Run("Payments", testPaymentsReloadAll)
	t.Run("Students", testStudentsReloadAll)
	t.Run("Subjects", testSubjectsReloadAll)
	t.Run("Teachers", testTeachersReloadAll)
	t.Run("Teaches", testTeachesReloadAll)
	t.Run("Timetables", testTimetablesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Accounts", testAccountsSelect)
	t.Run("Admins", testAdminsSelect)
	t.Run("Appointments", testAppointmentsSelect)
	t.Run("Classes", testClassesSelect)
	t.Run("Grades", testGradesSelect)
	t.Run("Notifications", testNotificationsSelect)
	t.Run("Parents", testParentsSelect)
	t.Run("Payments", testPaymentsSelect)
	t.Run("Students", testStudentsSelect)
	t.Run("Subjects", testSubjectsSelect)
	t.Run("Teachers", testTeachersSelect)
	t.Run("Teaches", testTeachesSelect)
	t.Run("Timetables", testTimetablesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Accounts", testAccountsUpdate)
	t.Run("Admins", testAdminsUpdate)
	t.Run("Appointments", testAppointmentsUpdate)
	t.Run("Classes", testClassesUpdate)
	t.Run("Grades", testGradesUpdate)
	t.Run("Notifications", testNotificationsUpdate)
	t.Run("Parents", testParentsUpdate)
	t.Run("Payments", testPaymentsUpdate)
	t.Run("Students", testStudentsUpdate)
	t.Run("Subjects", testSubjectsUpdate)
	t.Run("Teachers", testTeachersUpdate)
	t.Run("Teaches", testTeachesUpdate)
	t.Run("Timetables", testTimetablesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceUpdateAll)
	t.Run("Admins", testAdminsSliceUpdateAll)
	t.Run("Appointments", testAppointmentsSliceUpdateAll)
	t.Run("Classes", testClassesSliceUpdateAll)
	t.Run("Grades", testGradesSliceUpdateAll)
	t.Run("Notifications", testNotificationsSliceUpdateAll)
	t.Run("Parents", testParentsSliceUpdateAll)
	t.Run("Payments", testPaymentsSliceUpdateAll)
	t.Run("Students", testStudentsSliceUpdateAll)
	t.Run("Subjects", testSubjectsSliceUpdateAll)
	t.Run("Teachers", testTeachersSliceUpdateAll)
	t.Run("Teaches", testTeachesSliceUpdateAll)
	t.Run("Timetables", testTimetablesSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("Accounts", testAccountsUpsert)
	t.Run("Admins", testAdminsUpsert)
	t.Run("Appointments", testAppointmentsUpsert)
	t.Run("Classes", testClassesUpsert)
	t.Run("Grades", testGradesUpsert)
	t.Run("Notifications", testNotificationsUpsert)
	t.Run("Parents", testParentsUpsert)
	t.Run("Payments", testPaymentsUpsert)
	t.Run("Students", testStudentsUpsert)
	t.Run("Subjects", testSubjectsUpsert)
	t.Run("Teachers", testTeachersUpsert)
	t.Run("Teaches", testTeachesUpsert)
	t.Run("Timetables", testTimetablesUpsert)
}
