package repository

import (
	"database/sql"
	"github.com/middleware2018-PSS/Services/src/models"
)

func (r *postgresRepository) UpdateParent(id int64) (err error) {
	// TODO
	return
}

func (r *postgresRepository) ParentById(id int64) (interface{}, error) {
	p := models.Parent{}
	err := r.QueryRow(`SELECT id,	name, surname, mail, info
								FROM back2school.parents WHERE id = $1`,
		id).Scan(&p.ID, &p.Name, &p.Surname, &p.Mail, &p.Info)
	return p, switchError(err)
}

func (r *postgresRepository) ChildrenByParent(id int64, offset int, limit int) (children []interface{}, err error) {
	return r.listByID(id, offset, limit, `SELECT s.id, s.name, s.surname, s.mail, s.info
								FROM back2school.isparent join back2school.students as s on student = s.id 
								WHERE parent = $1
								order by s.name desc
								limit $2 offset $3`,
		func(rows *sql.Rows) (interface{}, error) {
			student := models.Student{}
			err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student, err
		})
}

func (r *postgresRepository) PaymentsByParent(id int64, offset int, limit int) (payments []interface{}, err error) {
	return r.listByID(id, offset, limit, `select p.id, p.amount, p.student, p.payed, p.reason, p.emitted
		from back2school.payments as p natural join back2school.isparent
		where parent = $1
		order by p.emitted desc
		limit $2 offset $3`, func(rows *sql.Rows) (interface{}, error) {
		payment := models.Payment{}
		err := rows.Scan(&payment.ID, &payment.Amount, &payment.Student.ID, &payment.Payed, &payment.Reason, &payment.Emitted)
		return payment, err
	})
}

func (r *postgresRepository) NotificationsByParent(id int64, offset int, limit int) (list []interface{}, err error) {
	return r.listByID(id, offset, limit, `select * from (
				select n.id, n.receiver, n.message, n.receiver_kind, n.time
				from back2school.notification as n join back2school.isparent on n.receiver = student
				where parent = $1 and receiver_kind = 'student'
				union all 
				select n.id, n.receiver, n.message, n.receiver_kind, n.time
				from back2school.notification as n
				where receiver_kind = 'general'
				) as a order by time desc
				limit $2 offset $3`,
		func(rows *sql.Rows) (interface{}, error) {
			notification := models.Notification{}
			err := rows.Scan(&notification.ID, &notification.Receiver, &notification.Message,
				&notification.ReceiverKind, &notification.Time)
			return notification, err
		})
}

func (r *postgresRepository) AppointmentsByParent(id int64, offset int, limit int) (appointments []interface{}, err error) {
	return r.listByID(id, offset, limit, `select a.id, a.student, a.teacher, a.location, a.time
				from back2school.appointments as a natural join back2school.isparent 
				where parent = $1
				order by a.time desc
				limit $2 offset $3`, func(rows *sql.Rows) (interface{}, error) {
		appointment := models.Appointment{}
		err := rows.Scan(&appointment.ID, &appointment.Student, &appointment.Teacher, &appointment.Location, &appointment.Time)
		return appointment, err
	})
}
