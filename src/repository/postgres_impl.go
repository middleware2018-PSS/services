package repository

import (
	"database/sql"
	_ "github.com/middleware2018-PSS/Services/src/docs"
	"github.com/middleware2018-PSS/Services/src/models"
)

type Subjects struct {
	Subjects []string `json:"subjects" example:"science"`
}

func (r *postgresRepository) CheckUser(userID string, password string) (int, string, bool) {
	query := `select id, kind from back2school.accounts where "user" = $1 and "password" = $2`
	var id int
	var kind string
	err := r.QueryRow(query, userID, password).Scan(&id, &kind)
	return id, kind, err == nil
}

// @Summary Get a appointment by id
// @Param id path int true "Appointment ID"
// @Tags Appointments
// @Success 200 {object} models.Appointment
// @Router /appointments/{id} [get]
func (r *postgresRepository) AppointmentByID(id int, who int, whoKind string) (interface{}, error) {
	appointment := models.Appointment{}
	var args []interface{}
	var query string
	switch whoKind {
	case ParentUser:
		query = "SELECT id, student, teacher, time, location " +
			"FROM back2school.appointments natural join back2school.isParent WHERE id = $1 and parent = $2"
		args = append(args, id, who)
	case TeacherUser:
		query = "SELECT id, student, teacher, time, location " +
			"FROM back2school.appointments WHERE id = $1 and teacher = $2"
		args = append(args, id, who)
	case AdminUser:
		query = "SELECT id, student, teacher, time, location " +
			"FROM back2school.appointments WHERE id = $1 "
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query, id, who).Scan(
		&appointment.ID, &appointment.Student.ID, &appointment.Teacher.ID, &appointment.Time, &appointment.Location)
	return switchResult(appointment, err)
}

// @Summary Get a grade by id
// @Param id path int true "Grade ID"
// @Tags Grades
// @Success 200 {object} models.Grade
// @Router /grades/{id} [get]
func (r *postgresRepository) GradeByID(id int, who int, whoKind string) (interface{}, error) {
	grade := models.Grade{}
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "SELECT id, student, teacher, subject, date, grade " +
			"FROM back2school.grades natural join back2school.isParent WHERE id = $1 and parent = $2 "
		args = append(args, id, who)
	case TeacherUser:
		query = "SELECT id, student, teacher, subject, date, grade " +
			"FROM back2school.grades WHERE id = $1 and teacher = $2"
		args = append(args, id, who)
	case AdminUser:
		query = "SELECT id, student, teacher, subject, date, grade " +
			"FROM back2school.grades WHERE id = $1"
	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query, id, who).Scan(
		&grade.ID, &grade.Student.ID, &grade.Teacher.ID, &grade.Subject, &grade.Date, &grade.Grade)
	return switchResult(grade, err)
}

// @Summary Get a class by id
// @Param id path int true "Class ID"
// @Tags Classes
// @Success 200 {object} models.Class
// @Router /classes/{id} [get]
func (r *postgresRepository) ClassByID(id int, who int, whoKind string) (interface{}, error) {
	class := models.Class{}
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		query = "SELECT id, year, section, info, grade FROM back2school.classes join back2school.teaches on class = id" +
			"WHERE id = $1 and teacher = $2"
		args = append(args, id, who)
	case AdminUser:
		query = "SELECT id, year, section, info, grade FROM back2school.classes " +
			"WHERE id = $1"
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query, args...).Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
	return switchResult(class, err)
}

// @Summary Get all classes
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Classes
// @Success 200 {array} models.Class
// @Router /classes [get]
func (r *postgresRepository) Classes(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		query = "select id, year, section, info, grade " +
			"from back2school.classes join back2school.teaches on class = id" +
			"WHERE teacher = $1" +
			"order by year desc, grade asc, section asc"
		args = append(args, who)
	case AdminUser:
		query = "select id, year, section, info, grade " +
			"from back2school.classes " +
			"order by year desc, grade asc, section asc"
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		class := models.Class{}
		err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
		return class, err
	}, limit, offset, args...)
}

// @Summary Get a student by class
// @Param id path int true "Class ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Classes
// @Success 200 {array} models.Student
// @Router /classes/{id}/students [get]
func (r *postgresRepository) StudentsByClass(id int, limit int, offset int, who int, whoKind string) (students []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		query = "select distinct s.id, s.name, s.surname, s.mail, s.info " +
			"from back2school.students as s join back2school.enrolled as e " +
			" join back2school.teaches as t on s.id = e.student and t.class = e.class" +
			"where s.class = $1 and t.teacher = $2" +
			"order by s.name desc, s.surname desc"
		args = append(args, id, who)
	case AdminUser:
		query = "select distinct id, name, surname, mail, info " +
			"from back2school.students join back2school.enrolled on student = id " +
			"where class = $1 " +
			"order by name desc, surname desc"
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			student := models.Student{}
			err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student, err
		}, limit, offset, args...)
}

//TODO add entrypoint !!!
// @Summary Get a lecture by class
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Param id path int true "Class ID"
// @Tags Lectures
// @Success 200 {array} models.Appointment
// @Router /classes/{id}/lectures [get]
func (r *postgresRepository) LectureByClass(id int, limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		query = "select id, class, subject, \"start\", \"end\", location, info " +
			"from back2school.timetable natural join back2school.teaches " +
			"where teacher = $1 and class = $2" +
			"order by \"start\" desc"
		args = append(args, who, id)
	case AdminUser:
		query = "select id, class, subject, \"start\", \"end\", location, info " +
			"from back2school.timetable" +
			"order by \"start\" desc"
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(
		query,
		func(rows *sql.Rows) (interface{}, error) {
			lecture := models.TimeTable{}
			err := rows.Scan(&lecture.ID, &lecture.Class, &lecture.Subject, &lecture.Start, &lecture.End, &lecture.Location, &lecture.Info)
			return lecture, err
		}, limit, offset, args...)
}

// @Summary Get a notification by id
// @Param id path int true "Notification ID"
// @Tags Notifications
// @Success 200 {object} models.Notification
// @Router /notifications/{id} [get]
func (r *postgresRepository) NotificationByID(id int, who int, whoKind string) (interface{}, error) {
	n := models.Notification{}
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		query = "SELECT id, receiver, message, time, receiver_kind " +
			"FROM back2school.notification" +
			"WHERE id = $1 and receiver = $2 and receiver_kind = $3"
		args = append(args, id, who, whoKind)
	case ParentUser:
		query = "SELECT id, receiver, message, time, receiver_kind " +
			"FROM back2school.notification WHERE id = $1 and receiver = $2 and receiver_kind = $3"
		args = append(args, id, who, whoKind)
	case AdminUser:
		query = "SELECT id, receiver, message, time, receiver_kind " +
			"FROM back2school.notification WHERE id = $1"
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query, args...).Scan(&n.ID,
		&n.Receiver, &n.Message, &n.Time, &n.ReceiverKind)
	return switchResult(n, err)
}

// List all notifications
// @Summary Get all notifications
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Notifications
// @Success 200 {array} models.Notification
// @Router /notifications [get]
func (r *postgresRepository) Notifications(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser, ParentUser:
		query = "select id, receiver, message, time, receiver_kind " +
			"from back2school.notification " +
			" where receiver = $1 and receiver_kind = $2" +
			"order by time desc, receiver_kind desc"
		args = append(args, who, whoKind)
	case AdminUser:
		query = "select id, receiver, message, time, receiver_kind " +
			"from back2school.notification " +
			"order by time desc, receiver_kind desc"
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			notification := models.Notification{}
			err := rows.Scan(&notification.ID, &notification.Receiver, &notification.Message, &notification.Time, &notification.ReceiverKind)
			return notification, err
		}, limit, offset, args...)
}

// Parents
// see/modify their personal data
// @Summary Get a parent by id
// @Param id path int true "Account ID"
// @Tags Parents
// @Success 200 {object} models.Parent
// @Router /parents/{id} [get]
func (r *postgresRepository) ParentByID(id int, who int, whoKind string) (interface{}, error) {
	p := models.Parent{}
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		if id == who {
			query = "SELECT id,	name, surname, mail, info " +
				"FROM back2school.parents WHERE id = $1"
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT id,	name, surname, mail, info " +
			"FROM back2school.parents WHERE id = $1"
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query,
		args).Scan(&p.ID, &p.Name, &p.Surname, &p.Mail, &p.Info)
	return switchResult(p, err)
}

// @Summary Get all grades
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Grades
// @Success 200 {array} models.Grade
// @Router /grades [get]
func (r *postgresRepository) Grades(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "select id, student, grade, subject, date, teacher " +
			"from back2school.grades natural join back2school.isParent" +
			" where parent = $1" +
			"order by date desc, teacher asc"
		args = append(args, who)
	case AdminUser:
		query = "select id, student, grade, subject, date, teacher " +
			"from back2school.grades " +
			"order by date desc, teacher asc"
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			g := models.Grade{}
			err := rows.Scan(&g.ID, &g.Student.ID, &g.Grade, &g.Subject, &g.Date, &g.Teacher)
			return g, err
		}, limit, offset, args...)
}

// @Summary Get all parents
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Parent
// @Router /parents [get]
func (r *postgresRepository) Parents(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		p, err := r.ParentByID(who, who, whoKind)
		return []interface{}{p}, err
		args = append(args, who)
	case AdminUser:
		query = "select id, name, surname, mail, info " +
			"from back2school.parents " +
			"order by name desc, surname desc"
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			parent := models.Parent{}
			err := rows.Scan(&parent.ID, &parent.Name, &parent.Surname, &parent.Mail, &parent.Info)
			return parent, err
		}, limit, offset, args...)
}

// see/modify the personal data of their registered children
// @Summary Get children of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Student
// @Router /parents/{id}/students [get]
func (r *postgresRepository) ChildrenByParent(id int, limit int, offset int, who int, whoKind string) (children []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		if id == who {
			query = "SELECT distinct s.id, s.name, s.surname, s.mail, s.info " +
				"FROM back2school.isparent join back2school.students as s on student = s.id  " +
				"WHERE parent = $1 " +
				"order by s.name desc"
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT distinct s.id, s.name, s.surname, s.mail, s.info " +
			"FROM back2school.isparent join back2school.students as s on student = s.id  " +
			"WHERE parent = $1 " +
			"order by s.name desc"
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			student := models.Student{}
			err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student, err
		}, limit, offset, args...)
}

// see the monthly payments that have been made to the school in the past
// @Summary Get payments of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Payment
// @Router /parents/{id}/payments [get]
func (r *postgresRepository) PaymentsByParent(id int, limit int, offset int, who int, whoKind string) (payments []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		if id == who {
			query = "select p.id, p.amount, p.student, p.payed, p.reason, p.emitted " +
				"from back2school.payments as p natural join back2school.isparent " +
				"where parent = $1 " +
				"order by p.emitted desc"
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "select p.id, p.amount, p.student, p.payed, p.reason, p.emitted " +
			"from back2school.payments as p natural join back2school.isparent " +
			"where parent = $1 " +
			"order by p.emitted desc"
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		payment := models.Payment{}
		err := rows.Scan(&payment.ID, &payment.Amount, &payment.Student.ID, &payment.Payed, &payment.Reason, &payment.Emitted)
		return payment, err
	}, limit, offset, args...)
}

// see general/personal notifications coming from the school
// @Summary Get notifications of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Notification
// @Router /parents/{id}/notifications [get]
func (r *postgresRepository) NotificationsByParent(id int, limit int, offset int, who int, whoKind string) (list []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		if id == who {
			query = "select * from ( " +
				"select n.id, n.receiver, n.message, n.receiver_kind, n.time " +
				"from back2school.notification as n join back2school.isparent on n.receiver = student " +
				"where parent = $1 and receiver_kind = 'student' " +
				"union all  " +
				"select n.id, n.receiver, n.message, n.receiver_kind, n.time " +
				"from back2school.notification as n " +
				"where receiver_kind = 'general' " +
				") as a order by time desc"
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "select * from ( " +
			"select n.id, n.receiver, n.message, n.receiver_kind, n.time " +
			"from back2school.notification as n join back2school.isparent on n.receiver = student " +
			"where parent = $1 and receiver_kind = 'student' " +
			"union all  " +
			"select n.id, n.receiver, n.message, n.receiver_kind, n.time " +
			"from back2school.notification as n " +
			"where receiver_kind = 'general' " +
			") as a order by time desc"
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			notification := models.Notification{}
			err := rows.Scan(&notification.ID, &notification.Receiver, &notification.Message,
				&notification.ReceiverKind, &notification.Time)
			return notification, err
		}, limit, offset, args...)
}

// see/modify appointments that they have with their children's teachers
// (calendar-like support for requesting appointments, err error)
// @Summary Get appointments of the parent
// @Param id path int true "Parent ID"
// @Tags Parents
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Appointment
// @Router /parents/{id}/appointments [get]
func (r *postgresRepository) AppointmentsByParent(id int, limit int, offset int, who int, whoKind string) (appointments []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		if id == who {
			query = "select a.id, a.student, a.teacher, a.location, a.time " +
				"from back2school.appointments as a natural join back2school.isparent  " +
				"where parent = $1 " +
				"order by a.time desc"
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "select a.id, a.student, a.teacher, a.location, a.time " +
			"from back2school.appointments as a natural join back2school.isparent  " +
			"where parent = $1 " +
			"order by a.time desc"
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			appointment := models.Appointment{}
			err := rows.Scan(&appointment.ID, &appointment.Student, &appointment.Teacher, &appointment.Location, &appointment.Time)
			return appointment, err
		}, limit, offset, args...)
}

// Get payment by id
// @Summary Get a payment by id
// @Param id path int true "Payment ID"
// @Tags Payments
// @Success 200 {object} models.Payment
// @Router /payments/{id} [get]
func (r *postgresRepository) PaymentByID(id int, who int, whoKind string) (interface{}, error) {
	payment := &models.Payment{}
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "SELECT id, amount, payed, emitted, reason " +
			"FROM back2school.payments natural join back2school.isParent" +
			" WHERE id = $1 and parent = $2"
		args = append(args, id, who)
	case AdminUser:
		query = "SELECT id, amount, payed, emitted, reason " +
			"FROM back2school.payments WHERE id = $1 "
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query, args...).Scan(payment.ID, payment.Amount, payment.Payed, payment.Emitted, payment.Reason)
	return switchResult(payment, err)
}

// @Summary Get all payments
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Payments
// @Success 200 {array} models.Payment
// @Router /payments [get]
func (r *postgresRepository) Payments(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "select id, amount, student, payed, reason, emitted " +
			"from back2school.payments natural join back2school.isParent " +
			" where parent = $1" +
			"order by payed asc, emitted asc"
		args = append(args, who)
	case AdminUser:
		query = "select id, amount, student, payed, reason, emitted " +
			"from back2school.payments " +
			"order by payed asc, emitted asc"
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			payment := models.Payment{}
			err := rows.Scan(&payment.ID, &payment.Amount, &payment.Student.ID, &payment.Payed, &payment.Reason, &payment.Emitted)
			return payment, err
		}, limit, offset, args)
}

// @Summary Get all appointments
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Appointments
// @Router /appointments [get]
// @Success 200 {object} representations.List
// @Security ApiKeyAuth
func (r *postgresRepository) Appointments(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "select id, student, teacher, location, time " +
			"from back2school.appointments natural join back2school.isParent " +
			" where parent = $1" +
			"order by time desc, teacher asc"
		args = append(args, who)
	case AdminUser:
		query = "select id, student, teacher, location, time " +
			"from back2school.appointments " +
			"order by time desc, teacher asc"
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			appointment := models.Appointment{}
			err := rows.Scan(&appointment.ID, &appointment.Student.ID, &appointment.Teacher.ID, &appointment.Location, &appointment.Time)
			return appointment, err
		}, limit, offset)
}

// LectureByClass(id int, limit int, offset int) (students []interface{}, err error)
// TODO GradeStudent(grade models.Grade) error
// TODO
// parents:
// see/pay (fake payment) upcoming scheduled payments (monthly, material, trips, err error)
// admins:
// everything
// @Summary Get all students
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Students
// @Success 200 {array} models.Student
// @Router /students [get]
func (r *postgresRepository) Students(limit int, offset int, who int, whoKind string) (student []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "select id, name, surname, mail, info  " +
			"from back2school.students join back2school.isParent on id=student " +
			"where parent = $1" +
			"order by name desc, surname desc"
		args = append(args, who)
	case AdminUser:
		query = "select id, name, surname, mail, info  " +
			"from back2school.students " +
			"order by name desc, surname desc"
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		student := models.Student{}
		err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
		return student, err
	}, limit, offset, args...)
}

// Get student by id
// @Summary Get a student by id
// @Param id path int true "Student ID"
// @Tags Students
// @Success 200 {object} models.Student
// @Router /students/{id} [get]
func (r *postgresRepository) StudentByID(id int, who int, whoKind string) (student interface{}, err error) {
	s := models.Student{}
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "SELECT id,	name, surname, mail, info  " +
			"FROM back2school.students join back2school.isParent on student = id" +
			"WHERE id = $1 and parent = $2"
		args = append(args, id, who)
	case AdminUser:
		query = "SELECT id,	name, surname, mail, info  " +
			"FROM back2school.students WHERE id = $1"
	default:
		return nil, ErrorNotAuthorized
	}
	err = r.QueryRow(query, args...).Scan(&s.ID,
		&s.Name, &s.Surname, &s.Mail, &s.Info)
	return switchResult(s, err)
}

// see the grades obtained by their children
// @Summary Get grades of the student
// @Param id path int true "Student ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Students
// @Success 200 {array} models.Grade
// @Router /students/{id}/grades [get]
func (r *postgresRepository) GradesByStudent(id int, limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "SELECT id, student, subject, date, grade, teacher " +
			"FROM back2school.grades natural join back2school.isParent " +
			"WHERE student = $1 and parent = $2" +
			"order by date desc"
		args = append(args, id, who)
		//TODO case ParentUser: //GradesByStudent
	case AdminUser:
		query = "SELECT id, student, subject, date, grade, teacher " +
			"FROM back2school.grades " +
			"WHERE student = $1 " +
			"order by date desc"
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			grade := models.Grade{}
			err := rows.Scan(&grade.ID, &grade.Student.ID, &grade.Subject, &grade.Date, &grade.Grade, &grade.Teacher.ID)
			return grade, err
		}, limit, offset, args...)
}

// see/modify their personal data
// Get teacher by id
// @Summary Get a teacher by id
// @Param id path int true "Teacher ID"
// @Tags Teachers
// @Success 200 {object} models.Teacher
// @Router /teachers/{id} [get]
func (r *postgresRepository) TeacherByID(id int, who int, whoKind string) (interface{}, error) {
	teacher := models.Teacher{}
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if id == who {
			query = "SELECT id, name, surname, mail  " +
				"FROM back2school.teachers" +
				"WHERE id = $1"
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT id, name, surname, mail  " +
			"FROM back2school.teachers " +
			"WHERE id = $1"
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query,
		id).Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Mail)
	return switchResult(teacher, err)

}

// List all teachers
// @Summary Get all teachers
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.Teacher
// @Router /teachers [get]
func (r *postgresRepository) Teachers(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		t, err := r.TeacherByID(who, who, whoKind)
		return []interface{}{t}, err
	case AdminUser:
		query = "select id, name, surname, mail, info  " +
			"from back2school.teachers " +
			"order by name desc, surname desc"
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		teacher := models.Teacher{}
		err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Mail, &teacher.Info)
		return teacher, err
	}, limit, offset, args...)
}

// @Summary Get appointments of the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.Appointment
// @Router /teachers/{id}/appointments [get]
func (r *postgresRepository) AppointmentsByTeacher(id int, limit int, offset int, who int, whoKind string) (appointments []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if id == who {
			query = "SELECT id, student, teacher, location, time " +
				"FROM back2school.appointments  " +
				"WHERE teacher = $1 " +
				"order by time desc"
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT id, student, teacher, location, time " +
			"FROM back2school.appointments  " +
			"WHERE teacher = $1 " +
			"order by time desc"
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			app := models.Appointment{}
			err := rows.Scan(&app.ID, &app.Student.ID, &app.Teacher.ID, &app.Location, &app.Time)
			return app, err
		}, limit, offset, args...)
}

// @Summary Get notifications of the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.TimeTable
// @Router /teachers/{id}/notifications [get]
func (r *postgresRepository) NotificationsByTeacher(id int, limit int, offset int, who int, whoKind string) (notifications []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if id == who {
			query = "SELECT id, receiver, message, receiver_kind, time  " +
				"FROM back2school.notification " +
				"WHERE (receiver = $1 and receiver_kind = $2) or receiver_kind = 'general' " +
				"order by time desc"
			args = append(args, who, whoKind)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT id, receiver, message, receiver_kind, time  " +
			"FROM back2school.notification  " +
			"WHERE (receiver = $1 and receiver_kind = '" + TeacherUser + "') or receiver_kind = 'general' " +
			"order by time desc"
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			notif := models.Notification{}
			err := rows.Scan(&notif.ID, &notif.Receiver, &notif.Message, &notif.ReceiverKind, &notif.Time)
			return notif, err
		}, limit, offset, args...)
}

// @Summary Get subject taught by the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Success 200 {object} repository.Subjects
// @Tags Teachers
// @Router /teachers/{id}/subjects [get]
func (r *postgresRepository) SubjectsByTeacher(id int, limit int, offset int, who int, whoKind string) (notifications []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if id == who {
			query = "SELECT DISTINCT subject FROM back2school.teaches where teacher = $1 order by subject"
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT DISTINCT subject FROM back2school.teaches where teacher = $1 order by subject"
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			subj := ""
			err := rows.Scan(&subj)
			return subj, err
		}, limit, offset, args...)
}

// @Summary Get classes in which the subject is taught by the teacher
// @Param id path int true "Teacher ID"
// @Param subject path int true "Subject ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Success 200 {array} models.Class
// @Tags Teachers
// @Router /teachers/{id}/subjects/{subject} [get]
func (r *postgresRepository) ClassesBySubjectAndTeacher(teacher int, subject string, limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if teacher == who {
			query = "SELECT id, year, section, info, grade " +
				"FROM back2school.teaches join back2school.classes on id = class  " +
				"WHERE teacher = $1 and subject = $2 " +
				"order by year desc, grade asc, section desc "
			args = append(args, who, subject)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT id, year, section, info, grade " +
			"FROM back2school.teaches join back2school.classes on id = class  " +
			"WHERE teacher = $1 and subject = $2 " +
			"order by year desc, grade asc, section desc "
		args = append(args, teacher, subject)
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			class := models.Class{}
			err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
			return class, err
		}, limit, offset, args...)
}

// @Summary Get lectures taught by the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.TimeTable
// @Router /teachers/{id}/lectures [get]
func (r *postgresRepository) LecturesByTeacher(id int, limit int, offset int, who int, whoKind string) (lectures []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if id == who {
			query = "SELECT id, class, subject, location, start, \"end\", info	 " +
				"from back2school.timetable natural join back2school.teaches as t " +
				"where t.teacher = $1 " +
				"order by start desc"
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT id, class, subject, location, start, \"end\", info	 " +
			"from back2school.timetable natural join back2school.teaches as t " +
			"where t.teacher = $1 " +
			"order by start desc"
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			lecture := models.TimeTable{}
			err := rows.Scan(&lecture.ID, &lecture.Class.ID, &lecture.Subject, &lecture.Location, &lecture.Start, &lecture.End, &lecture.Info)
			return lecture, err

		}, limit, offset, args...)
}

// @Summary Get classes in which the teacher teaches
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.Class
// @Router /teachers/{id}/classes [get]
func (r *postgresRepository) ClassesByTeacher(id int, limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	type Class struct {
		models.Class
		Subject models.Subject `json:"subject,omitempty"`
	}
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if id == who {
			query = "SELECT id, year, section, info, grade, subject " +
				"FROM back2school.teaches join back2school.classes on id = class  " +
				"WHERE teacher = $1 " +
				"order by subject asc, year desc, grade asc, section desc "
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT id, year, section, info, grade, subject " +
			"FROM back2school.teaches join back2school.classes on id = class  " +
			"WHERE teacher = $1 " +
			"order by subject asc, year desc, grade asc, section desc "
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			class := Class{}
			err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade, &class.Subject)
			return class, err
		}, limit, offset, args...)
}

// @Summary Update teacher's data
// @Param id path int true "Teacher ID"
// @Param teacher body models.Teacher true "data"
// @Tags Teachers
// @Success 204 {object} models.Teacher
// @Router /teachers/{id} [put]
func (r *postgresRepository) UpdateTeacher(teacher models.Teacher, who int, whoKind string) (err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if teacher.ID == who {
			query = "UPDATE back2school.teachers" +
				" SET name = $1, surname = $2, mail = $3, info = $4 " +
				"where id = $5"
			args = append(args, teacher.Name, teacher.Surname, teacher.Mail, teacher.Info, who)
		} else {
			return ErrorNotAuthorized
		}
	case AdminUser:
		query = "UPDATE back2school.teachers" +
			" SET name = $1, surname = $2, mail = $3, info = $4 " +
			"where id = $5"
		args = append(args, teacher.Name, teacher.Surname, teacher.Mail, teacher.Info, teacher.ID)
	default:
		return ErrorNotAuthorized
	}
	return r.execUpdate(query, args...)
}

// @Summary Update parents's data
// @Param id path int true "Parent ID"
// @Param parent body models.Parent true "data"
// @Tags Parents
// @Success 201 {object} models.Parent
// @Router /parents/{id} [put]
func (r *postgresRepository) UpdateParent(parent models.Parent, who int, whoKind string) (err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		if parent.ID == who {
			query = "UPDATE back2school.parents" +
				" SET name = $1, surname = $2, mail = $3, info = $4 " +
				"where id = $5"
			args = append(args, parent.Name, parent.Surname, parent.Mail, parent.Info, who)
		} else {
			return ErrorNotAuthorized
		}
	case AdminUser:
		query = "UPDATE back2school.parents" +
			" SET name = $1, surname = $2, mail = $3, info = $4 " +
			"where id = $5"
		args = append(args, parent.Name, parent.Surname, parent.Mail, parent.Info, parent.ID)
	default:
		return ErrorNotAuthorized
	}

	return r.execUpdate(query, args...)
}

// @Summary Update student's data
// @Param id path int true "Student ID"
// @Param student body models.Student true "data"
// @Tags Students
// @Success 201 {object} models.Student
// @Router /students/{id} [put]
func (r *postgresRepository) UpdateStudent(student models.Student, who int, whoKind string) (err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "IF $1 in (select parent from back2school.isParent where student = $2) UPDATE back2school.student" +
			" SET name = $3, surname = $4, mail = $5, info = $6 " +
			"where id = $7"
		args = append(args, who, student.ID, student.Name, student.Surname, student.Mail, student.Info, student.ID)
	case AdminUser:
		query = "UPDATE back2school.student" +
			" SET name = $1, surname = $2, mail = $3, info = $4 " +
			"where id = $5"
		args = append(args, student.Name, student.Surname, student.Mail, student.Info, student.ID)
	default:
		return ErrorNotAuthorized
	}
	return r.execUpdate(query, args...)
}

// @Summary Update appointment's data
// @Param id path int true "Appointment ID"
// @Param appointment body models.Appointment true "data"
// @Tags Appointments
// @Success 201 {object} models.Appointment
// @Router /appointments/{id} [put]
func (r *postgresRepository) UpdateAppointment(appointment models.Appointment, who int, whoKind string) (err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "if $1 in (select parent from back2school.isParent where student = $2) UPDATE back2school.appointments " +
			"SET student = $3, teacher = $4, location = $5, time = $6 where id = $7"
		args = append(args, who, appointment.Student.ID, appointment.Student.ID, appointment.Teacher.ID, appointment.Location, appointment.Time, appointment.ID)
	case TeacherUser:
		if appointment.Teacher.ID == who {
			query = "UPDATE back2school.appointments " +
				"SET student = $1, teacher = $2, location = $3, time = $4 where id = $5"
			args = append(args, appointment.Student.ID, who, appointment.Location, appointment.Time, appointment.ID)
		}
	case AdminUser:
		query = "UPDATE back2school.appointments " +
			"SET student = $1, teacher = $2, location = $3, time = $4 where id = $5"
		args = append(args, appointment.Student.ID, appointment.Teacher.ID, appointment.Location, appointment.Time, appointment.ID)
	default:
		return ErrorNotAuthorized
	}
	return r.execUpdate(query, args...)
}

// @Summary Create appointment
// @Param id path int true "Appointment ID"
// @Param appointment body models.Appointment true "data"
// @Tags Appointments
// @Router /appointments [post]
// @Success 201 {object} models.Appointment
func (r *postgresRepository) CreateAppointment(appointment models.Appointment, who int, whoKind string) (id int, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "if $1 in (select parent from back2school.isParent where student = $2) INSERT INTO back2school.appointments " +
			" (student, teacher, location, time) VALUES ($3, $4, $5, $6)"
		args = append(args, who, appointment.Student.ID, appointment.Teacher.ID, appointment.Location, appointment.Time)
	case TeacherUser:
		if appointment.Teacher.ID == who {
			query = "INSERT INTO back2school.appointments " +
				" (student, teacher, location, time) VALUES ($1, $2, $3, $4)"
			args = append(args, appointment.Student.ID, who, appointment.Location, appointment.Time)
		} else {
			return 0, ErrorNotAuthorized
		}
	case AdminUser:
		query = "INSERT INTO back2school.appointments " +
			" (student, teacher, location, time) VALUES ($1, $2, $3, $4)"
		args = append(args, appointment.Student.ID, appointment.Teacher.ID, appointment.Location, appointment.Time)
	default:
		return 0, ErrorNotAuthorized
	}
	return r.exec(query, args...)
}

// @Summary Create parent
// @Tags Parents
// @Param parent body models.Parent true "data"
// @Tags Parents
// @Success 201 {object} models.Parent
// @Router /parents [post]
func (r *postgresRepository) CreateParent(parent models.Parent, who int, whoKind string) (int, error) {
	if whoKind == AdminUser {
		query := "INSERT INTO back2school.parents " +
			"(name, surname, mail, info) VALUES ($1, $2, $3, $4)"
		return r.exec(query, parent.Name, parent.Surname, parent.Mail, parent.Info)
	} else {
		return 0, ErrorNotAuthorized
	}
}

// @Summary Create teacher
// @Param teacher body models.Teacher true "data"
// @Tags Teachers
// @Router /teachers [post]
// @Success 201 {object} models.Teacher
func (r *postgresRepository) CreateTeacher(teacher models.Teacher, who int, whoKind string) (int, error) {
	if whoKind == AdminUser {
		query := "INSERT INTO back2school.teachers" +
			" (name, surname, mail, info)" +
			" VALUES ($1, $2, $3, $4)"
		return r.exec(query, teacher.Name, teacher.Surname, teacher.Mail, teacher.Info)
	} else {
		return 0, ErrorNotAuthorized
	}

}

// @Summary Create student
// @Param student body models.Student true "data"
// @Tags Students
// @Router /students [post]
// @Success 201 {object} models.Student
func (r *postgresRepository) CreateStudent(student models.Student, who int, whoKind string) (int, error) {
	if whoKind == AdminUser {
		query := "INSERT INTO back2school.students" +
			" (name, surname, mail, info) " +
			" VALUES ($1, $2, $3, $4)"
		return r.exec(query, student.Name, student.Surname, student.Mail, student.Info)
	} else {
		return 0, ErrorNotAuthorized
	}

}

// @Summary Create class
// @Param class body models.Class true "data"
// @Tags Classes
// @Router /classes [post]
// @Success 201 {object} models.Class
func (r *postgresRepository) CreateClass(class models.Class, who int, whoKind string) (int, error) {
	if whoKind == AdminUser {
		query := "INSERT INTO back2school.classes" +
			" (year, section, info, grade) " +
			" VALUES ($1, $2, $3, $4)"
		return r.exec(query, class.Year, class.Section, class.Info, class.Grade)
	} else {
		return 0, ErrorNotAuthorized
	}

}

// @Summary Update Class's data
// @Param id path int true "Class ID"
// @Param parent body models.Class true "data"
// @Tags Classes
// @Success 201 {object} models.Class
// @Router /classes/{id} [put]
func (r *postgresRepository) UpdateClass(class models.Class, who int, whoKind string) (err error) {
	if whoKind == AdminUser {
		query := "UPDATE back2school.classes " +
			" SET year = $1, section = $2, info = $3, grade = $4" +
			" where id = $5"
		return r.execUpdate(query, class.Year, class.Section, class.Info, class.Grade, class.ID)
	} else {
		return ErrorNotAuthorized
	}

}

// @Summary Create notification
// @Param class body models.Notification true "data"
// @Tags Notifications
// @Router /notifications [post]
// @Success 201 {object} models.Notification
func (r *postgresRepository) CreateNotification(notification models.Notification, who int, whoKind string) (int, error) {
	if whoKind == AdminUser {
		query := "insert into back2school.classes " +
			" (receiver, message, time, receiver_kind) " +
			" VALUES ($1, $2, $3, $4) "
		return r.exec(query, notification.Receiver, notification.Message, notification.Time, notification.ReceiverKind)
	} else {
		return 0, ErrorNotAuthorized
	}

}

// @Summary Update notification
// @Param id path int true "Notification ID"
// @Param class body models.Notification true "data"
// @Tags Notifications
// @Router /notifications/{id} [put]
// @Success 201 {object} models.Notification
func (r *postgresRepository) UpdateNotification(notification models.Notification, who int, whoKind string) error {
	if whoKind == AdminUser {
		query := "UPDATE back2school.notification " +
			"SET receiver = $1, message = $2, time = $3, receiver_kind = $4" +
			" where id = $5"
		return r.execUpdate(query, notification.Receiver, notification.Message, notification.Time, notification.ReceiverKind, notification.ID)
	} else {
		return ErrorNotAuthorized
	}

}

// @Summary Create grade
// @Param class body models.Grade true "data"
// @Tags Grades
// @Router /grades [post]
// @Success 201 {object} models.Grade
func (r *postgresRepository) CreateGrade(grade models.Grade, who int, whoKind string) (int, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if grade.Teacher.ID == who {
			query = "if ($5, $1, $3) in " +
				"(select t.teacher, t.subject, e.student from back2school.teaches as t natural join back2school.enrolled as e ) " +
				" insert into back2school.grades " +
				" (student, grade, subject, date, teacher) " +
				" VALUES ($1, $2, $3, $4, $5) "
			args = append(args, grade.Student.ID, grade.Grade, grade.Subject, grade.Date, grade.Teacher.ID)
		} else {
			return 0, ErrorNotAuthorized
		}
	case AdminUser:
		query = "insert into back2school.grades " +
			" (student, grade, subject, date, teacher) " +
			" VALUES ($1, $2, $3, $4, $5) "
		args = append(args, grade.Student.ID, grade.Grade, grade.Subject, grade.Date, grade.Teacher.ID)
	default:
		return 0, ErrorNotAuthorized
	}

	return r.exec(query, args...)
}

// @Summary Update Grade
// @Param id path int true "Grade ID"
// @Param class body models.Grade true "data"
// @Tags Grades
// @Router /grades/{id} [put]
// @Success 201 {object} models.Grade
func (r *postgresRepository) UpdateGrade(grade models.Grade, who int, whoKind string) error {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if grade.Teacher.ID == who {
			query = "UPDATE back2school.grades " +
				"SET student = $1, grade = $2, subject = $3, date = $4, teacher = $5) " +
				" where id = $6"
			args = append(args, grade.Student.ID, grade.Grade, grade.Subject, grade.Date, grade.Teacher.ID, grade.ID)
		} else {
			return ErrorNotAuthorized
		}
	case AdminUser:
		query = "UPDATE back2school.grades " +
			"SET student = $1, grade = $2, subject = $3, date = $4, teacher = $5) " +
			" where id = $6"
		args = append(args, grade.Student.ID, grade.Grade, grade.Subject, grade.Date, grade.Teacher.ID)
	default:
		return ErrorNotAuthorized
	}
	return r.execUpdate(query, args...)
}

// @Summary Create payment
// @Param class body models.Payment true "data"
// @Tags Payments
// @Router /payments [post]
// @Success 201 {object} models.Payment
func (r *postgresRepository) CreatePayment(payment models.Payment, who int, whoKind string) (int, error) {
	if whoKind == AdminUser {
		query := "insert into back2school.payments " +
			" (amount, student, payed, reason, emitted) " +
			" VALUES ($1, $2, $3, $4, $5) "
		return r.exec(query, payment.Amount, payment.Student.ID, payment.Payed, payment.Reason, payment.Emitted)
	} else {
		return 0, ErrorNotAuthorized
	}
}

// @Summary Update payment
// @Param id path int true "Payment ID"
// @Param class body models.Payment true "data"
// @Tags Payments
// @Router /payments/{id} [put]
// @Success 201 {object} models.Payment
func (r *postgresRepository) UpdatePayment(payment models.Payment, who int, whoKind string) error {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		query = "if $7 in (select parent from back2school.isParent where student = $2) UPDATE back2school.grades " +
			"SET amount = $1, student = $2, payed = $3, reason = $4, emitted = $5" +
			" where id = $6"
		args = append(args, payment.Amount, payment.Student.ID, payment.Payed, payment.Reason, payment.Emitted, payment.ID, who)

	case AdminUser:
		query = "UPDATE back2school.grades " +
			"SET amount = $1, student = $2, payed = $3, reason = $4, emitted = $5" +
			" where id = $6"
		args = append(args, payment.Amount, payment.Student.ID, payment.Payed, payment.Reason, payment.Emitted, payment.ID)
	default:
		return ErrorNotAuthorized
	}
	return r.execUpdate(query, args...)
}
